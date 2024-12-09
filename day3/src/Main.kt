fun main() {
    val lines = readInput("input")
    val part1Result = part1(lines)
    println("Day 3 Part 1")
    println("Result = $part1Result")

    val part2Result = part2(lines)
    println("Day 3 Part 2")
    println("Result = $part2Result")
}

fun part1(lines: List<String>): Int {
    val pattern = Regex("mul\\(\\d{1,3},\\d{1,3}\\)")
    var sum = 0
    for (line in lines) {
        val matchResult = pattern.findAll(line)
        for (match in matchResult) {
            sum += convertAndMultiply(match.value)
        }
    }
    return sum
}

fun part2(lines: List<String>): Int {
    val mulPattern = Regex("mul\\(\\d{1,3},\\d{1,3}\\)")
    val enablePattern = Regex("do\\(\\)|don't\\(\\)")
    var enable = true
    var sum = 0
    for (line in lines) {
        val mulMatchResult = mulPattern.findAll(line)
        val enableMatchList = enablePattern.findAll(line).toMutableList()
        for (match in mulMatchResult) {
            val foundLogicOperator = enableMatchList.firstOrNull { it.range.last < match.range.first }
            if (foundLogicOperator != null) {
                enableMatchList.remove(foundLogicOperator)
                enable = foundLogicOperator.value == "do()"
            }
            if (!enable) {
                continue
            }
            sum += convertAndMultiply(match.value)
        }
    }
    return sum
}

private const val mulRegex = """mul\((\d{1,3}),(\d{1,3})\)"""
private const val doRegex = """do\(\)"""
private const val dontRegex = """don't\(\)"""

fun part1Fancy(lines: List<String>): Int {
    return lines.sumOf { line ->
        mulRegex.toRegex().findAll(line).sumOf {
            val (first: String, second: String) = it.destructured
            first.toInt() * second.toInt()
        }
    }
}

fun part2Fancy(lines: List<String>): Int {
    var sum = 0
    var enabled = true

    """$mulRegex|$doRegex|$dontRegex""".toRegex().findAll(lines.joinToString()).forEach { match ->
        when(match.value) {
            "don't()" -> enabled = false
            "do()" -> enabled = true
            else -> if (enabled) sum += match.multiplyNumbers()
        }
    }

    return sum
}

private fun MatchResult.multiplyNumbers(): Int {
    val (first: String, second: String) = destructured
    return first.toInt() * second.toInt()
}

fun convertAndMultiply(x: String): Int {
    val mulStr = x.replace("mul(", "").replace(")", "")
    val split = mulStr.split(",")
    return split[0].toInt() * split[1].toInt()
}