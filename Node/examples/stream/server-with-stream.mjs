import fs from 'node:fs'
import { createServer } from 'node:http'

const server = createServer()

server.on('request', (req, res) => {
    const src = fs.createReadStream('./big.file')

    src.pipe(res)
})

server.listen(8000)
