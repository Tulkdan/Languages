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

class WrittenNumber(
    val written: String,
    val size: Int,
    val toReplace: String
) {
    fun getIndexAtString(str: String): Int {
	return str.indexOf(written)
    }

    fun replaceAtString(str: String): String {
	return str.replace(written, toReplace)
    }

    override fun toString(): String {
	return "$written $size"
    }
}

fun checkIdxForEachWrittenNumber(line: String): MutableList<WrittenNumber> {
    return mutableListOf(
	WrittenNumber("one", 3, "1"),
	WrittenNumber("two", 3, "2"),
	WrittenNumber("three", 5, "3"),
	WrittenNumber("four", 4, "4"),
	WrittenNumber("five", 4, "5"),
	WrittenNumber("six", 3, "6"),
	WrittenNumber("seven", 5, "7"),
	WrittenNumber("eight", 5, "8"),
	WrittenNumber("nine", 4, "9")
    ).filter{ it.getIndexAtString(line) != -1 }
	.sortedBy { it.getIndexAtString(line) }
	.toMutableList()
}

fun main(args: Array<String>) {
    val filename = args.firstOrNull() ?: "./example.txt"

    val lines = openFile(filename)

    val result = lines.map { line -> run {
	val numbersWritten = checkIdxForEachWrittenNumber(line)

	for (index in 1..<(numbersWritten.size - 1)) {
	    val actualData = numbersWritten.getOrNull(index)

	    if (actualData == null) {
		continue
	    }

	    val idx = actualData.getIndexAtString(line)

	    val previousData = numbersWritten.getOrNull(index - 1)

	    if (previousData == null) {
		continue
	    }

	    val previousIdx = previousData.getIndexAtString(line)
	    val previousWordSize = previousData.size

	    if (idx in previousIdx..(previousIdx + previousWordSize)) {
		numbersWritten.removeAt(index)
	    }
	}

	val result = numbersWritten.fold(line) { acc, classOperator -> classOperator.replaceAtString(acc) }

	result.replace("""\D""".toRegex(), "")
    } }.fold(0) { acc, line -> run {
	val first = line.first().digitToInt()
	val last = line.last().digitToInt()

	acc + (first * 10) + last
    } }

    println("${result}")
}
