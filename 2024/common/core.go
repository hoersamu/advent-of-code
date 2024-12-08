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

func Solve[T comparable](part func(*bufio.Scanner) T) {
	file, scanner := openScanner("./input.txt")
	defer file.Close()
	fmt.Printf("Result: %v\n", part(scanner))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Test[T comparable](t *testing.T, expected T, part func(*bufio.Scanner) T) {
	TestWithPath(t, expected, part, "./example.txt")
}

func TestWithPath[T comparable](t *testing.T, expected T, part func(*bufio.Scanner) T, path string) {
	file, scanner := openScanner(path)
	defer file.Close()
	result := part(scanner)
	if expected != result {
		t.Errorf("Result is %v, not %v", result, expected)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
