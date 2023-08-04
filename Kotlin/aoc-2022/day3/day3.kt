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

    var myPoints = 0

    for (i in 0..(lines.count() - 1) step 3) {
	val first = lines.get(i).toList()
	val second = lines.get(i+1).toList()
	val third = lines.get(i+2).toList()

	val duplicate = first.intersect(second).intersect(third)

	val letter = duplicate.first()
	val value = points.get(letter) ?: 0

	myPoints += value
    }

    println(myPoints)
}
