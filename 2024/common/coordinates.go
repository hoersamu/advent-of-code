package common

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

// Returns the vector from the current point to the target point
func (p Point) GetVector(target Point) Point {
	return Point{X: target.X - p.X, Y: target.Y - p.Y}
}

// Returns the inverted vector (i.e. the vector pointing in the opposite direction)
func (p Point) Invert() Point {
	return Point{X: -p.X, Y: -p.Y}
}

// Returns the point that is the sum of the current point and the vector
func (p Point) Add(vector Point) Point {
	return Point{X: p.X + vector.X, Y: p.Y + vector.Y}
}

// Returns a string representation of the point
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// Represents a grid with a set size
type Grid struct {
	SizeX int
	SizeY int
}

// Returns true if the point is within the grid
func (g Grid) Validate(p Point) bool {
	return p.X >= 0 && p.X < g.SizeX && p.Y >= 0 && p.Y < g.SizeY
}
