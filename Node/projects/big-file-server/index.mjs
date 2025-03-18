import fs from 'node:fs'
import http from 'node:http'
import pump from 'pump'

const server = http.createServer((req, res) => {
    const stream = fs.createReadStream('big.file')
    pump(stream, res, done)
})

function done(err) {
    if (err) {
	return console.error('File was not fully streamed to the user', err)
    }
    console.log('File was fully streamed to the user')
}

server.listen(3000)
