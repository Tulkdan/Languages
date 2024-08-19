import { Duplex } from 'node:stream'

const inoutStream = new Duplex({
    write(chunk, encoding, callback) {
        console.log(chunk.toString())
        callback()
    }

    read(size) {
        this.push(String.fromCharCode(this.currentCharCode++))
        if (this.currentCharCode > 90 ) {
            this.push(null)
        }
    }
})

inoutStream.currentCharCode = 65

process.stding
    .pipe(inountStream)
    .pipe(process.stdout)
