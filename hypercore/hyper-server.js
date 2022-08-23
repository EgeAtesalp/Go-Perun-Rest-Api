const kappa = require('kappa-core')
const ram = require('random-access-memory')
const crypto = require('crypto');
const hyperswarm = require('hyperswarm')
const pump = require('pump')
const recordView = require('./record-view');
const level = require('level-mem');
const http = require('http');

const topic = crypto.createHash('sha256').update('jsconf-todos').digest()
const swarm = hyperswarm()
var db = [] // This db is passed onto Go

const myArgs = process.argv.slice(2);

const core = kappa(ram, { valueEncoding: 'json' })
core.use('records', recordView(level({ valueEncoding: "json" })))
core.writer('local', function (err, feed) {
    swarm.join(topic, { lookup: true, announce: true })
    swarm.on('connection', function (connection, info) {
        console.log("New peer!")
        pump(connection, core.replicate(info.client, { live: true }), connection)
    });

    db.push({ alias: myArgs[0], address: myArgs[1], port: myArgs[2] })

    feed.append({
        type: 'put',
        id: Date.now(),
        alias: myArgs[0],
        address: myArgs[1],
        port: myArgs[2],
    })
})

core.ready([], function () {
    // Listen for latest message. 
    core.api.records.on('batch', function (data) {
        for (let msg of data) {
            // Add latest entry to db array
            db.push(msg.value)
            console.log(db)
        }
    });
})

// http stuffs

const server = http.createServer((req, res) => {
    if (req.url === '/') {
        res.write('Hello World')
        console.log(core)
        res.end()
    }

    if (req.url === '/db') {
        dbString = JSON.stringify(db)
        res.write('{"users":' + dbString + '}')
        res.end()
    }
})

server.listen(3002);