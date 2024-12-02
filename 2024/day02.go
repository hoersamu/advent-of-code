package main

func DayTwo() (int, int) {
	solution1 := DayTwoPartOne()
	solution2 := DayTwoPartTwo()
	return solution1, solution2
}

func isLineSafe(line []int, isLineAscending bool, skipIndex int) bool {
	isSafe := true

	lineWithoutIndex := line
	if skipIndex != -1 {
		lineWithoutIndex = RemoveIndex(line, skipIndex)
	}

	for i := 0; i < len(lineWithoutIndex)-1; i++ {
		multiplier := 1
		if !isLineAscending {
			multiplier = -1
		}
		num1 := lineWithoutIndex[i] * multiplier
		num2 := lineWithoutIndex[i+1] * multiplier

		distance := GetDistance(num1, num2)

		if num1 > num2 || distance < 1 || distance > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func DayTwoPartOne() int {
	lines := ReadFileWithDelimiterAsInt(
		// "./day02.example.input.txt",
		"./day02.input.txt",
		" ",
	)

	safeCount := 0

	for _, line := range lines {
		isLineAscending := line[0] < line[1]

		isSafe := isLineSafe(line, isLineAscending, -1)

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func DayTwoPartTwo() int {
	lines := ReadFileWithDelimiterAsInt(
		// "./day02.example.input.txt",
		"./day02.input.txt",
		" ",
	)

	safeCount := 0

	for _, line := range lines {
		for i := -1; i < len(line); i++ {
			isSafe := isLineSafe(line, true, i)
			if isSafe {
				safeCount++
				break
			}
			isSafe = isLineSafe(line, false, i)
			if isSafe {
				safeCount++
				break
			}
		}
	}

	return safeCount
}
