import java.util.*

fun main(args: Array<String>) {
    val unsortedValues = readValuesFromConsole()
    val sortedValues = Sorter.sortByBubbleSort(unsortedValues)
    print(sortedValues)
}

private fun readValuesFromConsole(): List<Int> {
    val scanner = Scanner(System.`in`)
    val inputCount = scanner.nextInt()
    val values = arrayListOf<Int>()
    (0 until inputCount).forEach { _ ->
        val value = scanner.nextInt()
        values.add(value)
    }
    return values
}
