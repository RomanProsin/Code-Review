fun main() {
    val array: Array<Int> = arrayOf(3, 2, 1)
    array.forEach { print(it) }
    println()
    val sortedArray = array.sort()
    sortedArray.forEach { print(it) }
}

private fun Array<Int>.sort(): Array<Int> {
    var temp: Int
    var min = 0
    for (i in this.indices) {
        for (j in i..(this.size - i - 2)) {
            if (this[j] > this[j + 1]) {
                min = this[j + 1]
                temp = this[j]
                this[j] = this[j + 1]
                this[j + 1] = temp
            }
        }
    }
    if (this[min] < this[0]) {
        temp = this[min]
        this[min] = this[0]
        this[0] = temp
    }
    return this
}
