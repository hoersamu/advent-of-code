package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func openScanner(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file, bufio.NewScanner(file)
}

func Solve(part func(*bufio.Scanner) string) {
	file, scanner := openScanner("./input.txt")
	defer file.Close()
	fmt.Printf("Result: %s\n", part(scanner))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Test(t *testing.T, expected string, part func(*bufio.Scanner) string) {
	TestWithPath(t, expected, part, "./example.txt")
}

func TestWithPath(t *testing.T, expected string, part func(*bufio.Scanner) string, path string) {
	file, scanner := openScanner(path)
	defer file.Close()
	result := part(scanner)
	if expected != result {
		t.Errorf("Result is %s, not %s", result, expected)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
