import fs from 'node:fs'
import { createServer } from 'node:http'

const server = createServer()

server.on('request', (req, res) => {
    fs.readFile('./big.file', (err, data) => {
        if (err) throw err;

        res.end(data)
    })
})

server.listen(8000)
