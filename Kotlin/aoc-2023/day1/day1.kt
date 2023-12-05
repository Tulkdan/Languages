import java.io.File
import java.io.IOException
import kotlin.system.exitProcess

fun openFile(filename: String): List<String> {
    try {
	return File(filename).readLines()
    } catch (e: IOException) {
	println("couldn't open file")
	exitProcess(1)
    }
}

fun main(args: Array<String>) {
    val filename = args.firstOrNull() ?: "./example.txt"

    val lines = openFile(filename)

    val result = lines.map {
	"""\D""".toRegex()
	    .replace(it, "")
    }.fold(0) { acc, line -> run {
	val first = line.first().digitToInt()
	val last = line.last().digitToInt()

	acc + (first * 10) + last
    } }

    println("${result}")
}
