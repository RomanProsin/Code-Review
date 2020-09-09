class Sorter {

    companion object {

        fun sortByBubbleSort(values: List<Int>): List<Int> {
            val valuesArray = values.toIntArray()
            for (index in 0 until (valuesArray.size - 1)) {
                for (currentPos in 0 until (valuesArray.size - index - 1)) {
                    if (valuesArray[currentPos] > valuesArray[currentPos + 1]) {
                        val tempValue = valuesArray[currentPos]
                        valuesArray[currentPos] = valuesArray[currentPos + 1]
                        valuesArray[currentPos + 1] = tempValue
                    }
                }
            }
            return valuesArray.toList()
        }
    }
}
