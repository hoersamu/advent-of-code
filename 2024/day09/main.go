package main

import (
	"bufio"
	"common"
	"day09/disk"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func createDiskFromScanner(scanner *bufio.Scanner) *disk.Disk {
	textStr := common.ScanToString(scanner)
	return disk.NewDisk(textStr)
}

func Part1(scanner *bufio.Scanner) int {
	disk := createDiskFromScanner(scanner)

	disk.DefragmentHard()
	checksum := disk.CalculateChecksum()
	return checksum
}

func Part2(scanner *bufio.Scanner) int {
	disk := createDiskFromScanner(scanner)

	disk.DefragmentSoft()
	checksum := disk.CalculateChecksum()
	return checksum
}
