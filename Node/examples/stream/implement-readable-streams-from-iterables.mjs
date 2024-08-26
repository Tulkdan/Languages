import { Readable } from 'node:stream'

async function* generate() {
    yield 'hello'
    yield 'streams'
}

const readable = Readable.from(generate())

readable.on('data', chunk => {
    console.log(chunk)
})
