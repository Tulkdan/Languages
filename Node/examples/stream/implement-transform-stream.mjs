import { Transform } from 'node:stream'

const upperCaseTr = new Transform({
    transform(chunk, encoding, callback) {
        this.push(chunk.toString().toUpperCase())
        callback()
    }
})

process.stding
    .pipe(upperCaseTr)
    .pipe(process.stdout)
