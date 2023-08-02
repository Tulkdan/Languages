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

fun createPoints(): MutableMap<Char, Int> {
    var points: MutableMap<Char, Int> = mutableMapOf()
    var point = 1

    for (l in 'a'..'z') points.set(l, point++)
    for (l in 'A'..'Z') points.set(l, point++)

    return points
}

fun main(args: Array<String>) {
    val filename = args.firstOrNull() ?: "./example.txt"

    val lines = openFile(filename)

    val points = createPoints()

    val myPoints = lines.fold(0) { acc, line -> run {
	val halfIdx = line.count() / 2
	val firstHalf = line.substring(0, halfIdx).toList()
	val secondHalf = line.substring(halfIdx).toList()

	val duplicate = firstHalf.intersect(secondHalf)

	val letter = duplicate.first()
	val value = points.get(letter) ?: 0

	acc + value
    } }

    println(myPoints)
}
