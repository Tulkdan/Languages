import fs from 'node:fs'
import { createServer } from 'node:http'
import { createGzip } from 'node:zlib'

const server = createServer()

server.on('request', (_, res) => {
  res.writeHead(200, { 'content-encoding': 'gzip' })

  fs.createReadStream('./big.file')
    .pipe(createGzip())
    .pipe(res)
})

server.listen(8000)
