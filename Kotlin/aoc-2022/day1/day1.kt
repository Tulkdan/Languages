import java.io.File

fun main() {
    val filename = "./elfs-calories-v2.txt"

    val lines: List<String> = File(filename).readLines()

    val elfs = lines.fold(mutableListOf(0)) { acc, line -> when(line) {
	"" -> {
	    acc.add(0)
	    acc
	}
	else -> {
	    val lastValue = acc.last()
	    acc[acc.count() - 1] = lastValue + line.toInt()
	    acc
	}
    } }

    val (first, second, third) = elfs.sortedDescending()

    println(first + second + third)
}
