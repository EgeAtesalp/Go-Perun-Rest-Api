#!/usr/bin/env python3

import argparse
from datetime import datetime
from time import sleep
import psutil
import threading
import json
import sys
from base import Measure

class Measurement:
    def __init__(self, print_to_terminal=True, log_path=None, log_type="json", 
        interval_in_ms=1000, duration_in_sec=None, cpu_stats=True, per_cpu=True, 
        cpu_meta=True, memory_stats=True, disk_stats=True, network_stats=True, 
        per_nic=False, sensor_stats=True):
        
        self.interval_in_ms = interval_in_ms
        self.duration_in_sec = duration_in_sec
        self.log_path = log_path
        self.log_type = log_type

        self.last_measure = None
        self.messures = 0
        self.print_to_terminal=print_to_terminal
        self.cpu_stats = cpu_stats
        self.per_cpu = per_cpu
        self.cpu_meta = cpu_meta
        self.memory_stats = memory_stats
        self.disk_stats = disk_stats
        self.network_stats = network_stats
        self.per_nic = per_nic
        self.sensor_stats = sensor_stats
        print("measurement have been prepared {0}\nto file: {1}".format(self,
            self.log_path))

    def startSingleMeasure(self):
        interval = self.interval_in_ms/1000
        self.timer = threading.Timer(interval, self.measure)
        self.timer.start()

    def startMeasurement(self):
        # enable timer
        self.started = datetime.now()
        self.busy = True
        self.startSingleMeasure()

    def stopOrRestartMeasurement(self):
        self.messures += 1
        # disable Timer
        if self.duration_in_sec != None:
            now = datetime.now()
            diff = now - self.started
            print("delta is {0}".format(diff.total_seconds()))
            if diff.total_seconds() > self.duration_in_sec:
                print("{1} measures after duration {0} sec exceeded".format(self.duration_in_sec, self.messures))
                self.busy = False
                return

        # restart        
        self.startSingleMeasure()

    def measure(self):
        currentMeasure = Measure()

        if self.cpu_stats:
            self.measure_cpu_load(currentMeasure)

        if self.memory_stats:
            self.measure_memory(currentMeasure)

        if self.network_stats:
            self.measure_network(currentMeasure)
        
        if self.disk_stats:
            self.measure_disk(currentMeasure)
        
        if self.sensor_stats:
            self.measure_sensors(currentMeasure)
        
        currentMeasure.finalize()
        self.last_measure = currentMeasure

        # restart timer as long not stopped
        self.stopOrRestartMeasurement()

        if self.log_path != None and self.log_path != "":
            self.print_results_to_path(currentMeasure)

        if self.print_to_terminal:
            self.print_results_to_terminal(currentMeasure)

    def print_results_to_terminal(self, currentMeasure):
        print("[{5}]{0}\n{1}\nduration in micsec: {2}\nfrom:{3}, until:{4}\n\n".format(datetime.now().strftime('%Y-%m-%dT%H:%M:%S.%fZ'), 
            currentMeasure.dumpToJsonStr(), currentMeasure.getDurationInMicSec(), currentMeasure.getCreated(), 
            currentMeasure.getFinalized(), self.messures))

    def print_results_to_path(self, currentMeasure):
        if self.log_type == "json":
            with open(self.log_path + '.' + self.log_type, 'a+') as log_file:
                log_file.write(currentMeasure.dumpToJsonStr())
                log_file.write('\n')
        else:
            print("log type [{0}] is not supported, yet.\n".format(self.log_type))

    def measure_cpu_info(self):
        data = {
            'cpu_use': psutil.Process().cpu_affinity(),
            'cpu_count_phys': psutil.cpu_count(),
            'cpu_count_logic': psutil.cpu_count(logical=True)
        }
        return data

    def measure_cpu_load(self, currentMeasure):
        data = {
            'cpu_sum': psutil.cpu_times()._asdict(),
            'cpu_freq_sum': psutil.cpu_freq()._asdict(),
            'cpu_load_history': [x / psutil.cpu_count() * 100 for x in psutil.getloadavg()]
        }

        if self.per_cpu:
            data['cpu_detail'] = psutil.cpu_times(percpu=True)._asdict()
            data['cpu_freq_detail'] = psutil.cpu_freq(percpu=True)._asdict()

        if self.cpu_meta:
            data['cpu_meta_info'] = self.measure_cpu_info()

        currentMeasure.update('cpu_load', data)

    def measure_memory(self, currentMeasure):
        data = {
            'ram_usage': psutil.virtual_memory()._asdict(),
            'swap_usage': psutil.swap_memory()._asdict(),
        }
        currentMeasure.update('memory_usage', data)

    def measure_disk(self, currentMeasure):
        data = {
            'disk_total': psutil.disk_usage('/')._asdict(),
            'disk_io': psutil.disk_io_counters(nowrap=True)._asdict(),
        }
        currentMeasure.update('disk_usage', data)

    def measure_network(self, currentMeasure):
        conns = []
        for con in psutil.net_connections():
            conAsDict = con._asdict()
            laddr = conAsDict['laddr']
            raddr = conAsDict['raddr']
            if len(laddr) > 0:
                conAsDict['laddr'] = conAsDict['laddr']._asdict()
            else:
                conAsDict['laddr'] = {}
            if len(raddr) > 0:
                conAsDict['raddr'] = conAsDict['raddr']._asdict()
            else:
                conAsDict['raddr'] = {}
            conns.append(conAsDict)

        data = {
            'network_connections': conns,
            'network_io_total': psutil.net_io_counters(nowrap=True)._asdict(),
        }
        if self.per_nic :
            data['network_io_per_nic'] = psutil.net_io_counters(nowrap=True, pernic=True)._asdict()
        currentMeasure.update('network_usage', data)

    def measure_sensors(self, currentMeasure):
        temp = psutil.sensors_temperatures()
        modTemps = []

        for cpuTermal in temp:
            key=0
            for shwtemp in temp[cpuTermal]:
                modTemp = {"label": shwtemp[0], "current": shwtemp[1], 
                        "high": shwtemp[2], "critical": shwtemp[3]}
                if(modTemp["label"] == ""):
                    modTemp["label"] = "t" + str(key)
                modTemps.append(modTemp)
                key+=1

        data = {
            'temperatures': modTemps
        }
        currentMeasure.update('sensor_data', data)

if __name__ == '__main__':
    p = argparse.ArgumentParser(
        description='tool for measure systems load while experiments'
    )

    p.add_argument('-f', '--file', required=True, default="measurement.json",
                   help='location to write the log file')

    p.add_argument('-t', '--type', required=True, default="json",
                   help='log file type e.g. json')
                   
    p.add_argument('--interval', help='interval in milliseconds', type=int, default=1000)
    p.add_argument('--duration', help='duration in seconds', type=int, default=None)

    p.add_argument('-v', '--verbose', help='enables output to the terminal', 
        action="store_true")
    
    p.add_argument

    parsed = p.parse_args()

    meas = Measurement(duration_in_sec=parsed.duration,
        interval_in_ms=parsed.interval,
        print_to_terminal=parsed.verbose,
        per_cpu=False,
        log_path=parsed.file,
        log_type=parsed.type)

    meas.startMeasurement()
    while(meas.busy):
        sleep(1)

    print("exit because measurement is done")
