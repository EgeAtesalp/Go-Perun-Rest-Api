from datetime import date, datetime
import json

class Measure:
    def __init__(self):
        self.finalized = False
        data = {
            'created': datetime.now()
        }
        self.data = data

    def getCreated(self):
        return self.data['created'].strftime('%Y-%m-%dT%H:%M:%S.%fZ')

    def getFinalized(self):
        return self.data['finalized'].strftime('%Y-%m-%dT%H:%M:%S.%fZ')

    def getDurationInMicSec(self):
        if self.data['created'] and self.data['finalized']:
            delta = self.data['finalized']-self.data['created']
            return delta.microseconds

    def update(self, name:str, value):
        if not self.finalized:
            self.data[name] = value
        else:
            raise AttributeError(name)

    def finalize(self):
        self.update('finalized', datetime.now())
        self.finalized = True

    def dumpToJsonStr(self):
        ser = self.data.copy()
        if ser['created']:
            ser['created'] = self.data['created'].strftime('%Y-%m-%dT%H:%M:%S.%fZ')
        if ser['finalized']:
            ser['finalized'] = self.data['finalized'].strftime('%Y-%m-%dT%H:%M:%S.%fZ')
        ser['duration_in_micsec'] = self.getDurationInMicSec()

        jDump = json.dumps(ser, default=lambda o: o.__dict__, sort_keys=True)
        return jDump