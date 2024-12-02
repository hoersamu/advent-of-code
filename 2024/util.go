package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func ReadFileWithDelimiter(filePath string, delimiter string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, delimiter)
		lines = append(lines, fields)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	return lines
}

func ReadFileWithDelimiterAsInt(filePath string, delimiter string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		for _, s := range strings.Split(line, delimiter) {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Unable to parse string to int", err)
			}
			nums = append(nums, num)
		}
		lines = append(lines, nums)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	return lines
}

func RemoveIndex(slice []int, s int) []int {
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:s]...)
	result = append(result, slice[s+1:]...)
	return result
}

func GetDistance(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}
