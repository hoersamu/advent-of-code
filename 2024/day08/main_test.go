package main

import (
	"common"
	"testing"
)

func Test1(t *testing.T) {
	common.Test(t, 14, Part1)
}

func Benchmark1(b *testing.B) {
	common.Solve(Part1)
}

func Test2(t *testing.T) {
	common.Test(t, 34, Part2)
}

func Benchmark2(b *testing.B) {
	common.Solve(Part2)
}
