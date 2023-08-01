import java.io.File

fun mapPoints(item: String) = when (item) {
    "X" -> 1
    "Y" -> 2
    "Z" -> 3
    else -> 0
}

/** Returns a list of if my play can [win, draw, lose] */
fun mapPlaysToWin(item: String) = when (item) {
    "X" -> listOf("C", "A", "B")
    "Y" -> listOf("A", "B", "C")
    "Z" -> listOf("B", "C", "A")
    "A" -> listOf("Z", "X", "Y")
    "B" -> listOf("X", "Y", "Z")
    "C" -> listOf("Y", "Z", "X")
    else -> listOf()
}

fun main(args: Array<String>) {
    val filename = args.firstOrNull() ?: "./example.txt"

    val lines: List<String> = File(filename).readLines()

    val myPoints = lines.fold(0) { acc, line -> run {
	val (enemyPlay, myPlay) = line.split(" ")

	val (win, draw, lose) = mapPlaysToWin(enemyPlay)

	val roundPoint = when (myPlay) {
	    "Z" -> 6 + mapPoints(lose)
	    "Y" -> 3 + mapPoints(draw)
	    else -> 0 + mapPoints(win)
	}

	acc + roundPoint
    } }

    println(myPoints)
}
