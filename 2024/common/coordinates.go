package common

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) GetVector(target Point) Point {
	return Point{X: target.X - p.X, Y: target.Y - p.Y}
}

func (p Point) Invert() Point {
	return Point{X: -p.X, Y: -p.Y}
}

func (p Point) Add(vector Point) Point {
	return Point{X: p.X + vector.X, Y: p.Y + vector.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Grid struct {
	SizeX int
	SizeY int
}

func (g Grid) Validate(p Point) bool {
	return p.X >= 0 && p.X < g.SizeX && p.Y >= 0 && p.Y < g.SizeY
}
