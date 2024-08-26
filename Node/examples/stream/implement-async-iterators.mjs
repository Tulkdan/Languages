import { createReadStream } from 'node:fs'

async function logChunks(readable) {
    for await (const chunk of readable) {
        console.log(chunk)
    }
}

const readable = craeateReadStream('./big.file', { encoding: 'utf8' })

logChunks(readable)
