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

    val points = lines.fold(0) { acc, line -> run {
	var (first, second) = line
	    .split(",")
	    .map {
		it.split("-")
		    .map { it.toInt() }
	    }

	println("${first.first()} <= ${second.first()} && ${first.last()} >= ${second.last()}")

	if (first.first() <= second.first() && first.last() >= second.last()) {
	    // first contains second
	    acc + 1
	} else if (first.first() >= second.first() && first.last() <= second.last()) {
	    // second contains
	    acc + 1
	} else acc
    } }

    println(points)
}
