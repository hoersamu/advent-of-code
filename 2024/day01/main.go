package main

import (
	"bufio"
	"common"
	"sort"
	"strconv"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) string {
	records := common.ScanWithDelimitersAsInt(scanner, "   ")

	var list1 []int
	var list2 []int
	for _, record := range records {
		list1 = append(list1, record[0])
		list2 = append(list2, record[1])
	}

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0

	for i, record := range list1 {
		totalDistance += common.GetDistance(record, list2[i])
	}

	return strconv.Itoa(totalDistance)
}

func Part2(scanner *bufio.Scanner) string {
	records := common.ScanWithDelimitersAsInt(scanner, "   ")
	var list1 []int
	list2 := make(map[int]int)

	for _, record := range records {
		list1 = append(list1, record[0])
		list2[record[1]] += 1
	}

	var similarityScore int = 0

	for _, num := range list1 {
		similarityScore += num * list2[num]
	}

	return strconv.Itoa(similarityScore)
}
