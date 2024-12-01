package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
)

func DayOne() {
	dayOnePartOne()
	dayOnePartTwo()
}

func readDayOneInput() [][]string {
	return ReadCsvFile("./day01.input.csv")
}

func dayOnePartOne() {
	records := readDayOneInput()
	var list1 []int
	var list2 []int
	for _, record := range records {
		num, err := strconv.Atoi(record[0])
		num2, err2 := strconv.Atoi(record[1])
		if err != nil || err2 != nil {
			log.Fatal("Unable to parse string to int", err, err2)
		}
		list1 = append(list1, num)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var distance int = 0

	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println("Day 1 Part 1:", distance)
}

func dayOnePartTwo() {
	records := readDayOneInput()
	var list1 []int
	list2 := make(map[int]int)
	for _, record := range records {
		num, err := strconv.Atoi(record[0])
		num2, err2 := strconv.Atoi(record[1])
		if err != nil || err2 != nil {
			log.Fatal("Unable to parse string to int", err, err2)
		}
		list1 = append(list1, num)
		list2[num2] += 1
	}

	var similarityScore int = 0

	for _, num := range list1 {
		similarityScore += num * list2[num]
	}

	fmt.Println("Day 1 Part 2:", similarityScore)
}

func DayOnePartOne() {
	records := ReadCsvFile("./01.input.csv")
	var list1 []int
	var list2 []int
	for _, record := range records {
		num, err := strconv.Atoi(record[0])
		num2, err2 := strconv.Atoi(record[1])
		if err != nil || err2 != nil {
			log.Fatal("Unable to parse string to int", err, err2)
		}
		list1 = append(list1, num)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var distance int = 0

	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println("Day 1 Part 1:", distance)
}
