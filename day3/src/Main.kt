fun main() {
    val lines = readInput("input")
    val part1Result = part1(lines)
    println("Day 3 Part 1")
    println("Result = $part1Result")
}

fun part1(lines: List<String>): Int {
    val pattern = Regex("mul\\(\\d{1,3},\\d{1,3}\\)")
    var sum = 0
    for (line in lines) {
        val matchResult = pattern.findAll(line)
        for (match in matchResult) {
            val mulStr = match.value.replace("mul(", "").replace(")", "")
            val split = mulStr.split(",")
            sum += split[0].toInt() * split[1].toInt()
        }
    }
    return sum
}