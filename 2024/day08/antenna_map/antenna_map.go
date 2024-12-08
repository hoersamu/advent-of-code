package antenna_map

import (
	"bufio"
	"common"
)

type AntennaMap struct {
	locations map[rune][]common.Point
	grid      common.Grid
	antiNodes map[int]map[int]bool
}

func NewAntennaMap(scanner *bufio.Scanner, useAllNodes bool) *AntennaMap {
	lines := common.ScanToLines(scanner)
	am := &AntennaMap{
		locations: parseLocations(lines),
		grid:      common.Grid{SizeX: len(lines[0]), SizeY: len(lines)},
		antiNodes: initializeAntiNodes(len(lines[0])),
	}
	am.calculateAntiNodes(useAllNodes)
	return am
}

func parseLocations(lines []string) map[rune][]common.Point {
	locations := make(map[rune][]common.Point)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				locations[char] = append(locations[char], common.Point{X: x, Y: y})
			}
		}
	}
	return locations
}

func initializeAntiNodes(width int) map[int]map[int]bool {
	antiNodes := make(map[int]map[int]bool)
	for x := 0; x < width; x++ {
		antiNodes[x] = make(map[int]bool)
	}
	return antiNodes
}

func (am *AntennaMap) calculateAntiNodes(useAllNodes bool) {
	for _, locations := range am.locations {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				if useAllNodes {
					am.addAllAntiNodes(locations[i], locations[j])
				} else {
					am.addDirectAntiNodes(locations[i], locations[j])
				}
			}
		}
	}
}

func (am *AntennaMap) addDirectAntiNodes(p1, p2 common.Point) {
	vector := p1.GetVector(p2)
	antiNode1 := p2.Add(vector)
	antiNode2 := p1.Add(vector.Invert())
	if am.grid.Validate(antiNode1) {
		am.antiNodes[antiNode1.X][antiNode1.Y] = true
	}
	if am.grid.Validate(antiNode2) {
		am.antiNodes[antiNode2.X][antiNode2.Y] = true
	}
}

func (am *AntennaMap) addAllAntiNodes(p1, p2 common.Point) {
	vector := p1.GetVector(p2)
	rightAntiNodes := []common.Point{p2}
	for {
		next := rightAntiNodes[len(rightAntiNodes)-1].Add(vector)
		if !am.grid.Validate(next) {
			break
		}
		rightAntiNodes = append(rightAntiNodes, next)
	}
	inverseVector := vector.Invert()
	leftAntiNodes := []common.Point{p1}
	for {
		next := leftAntiNodes[len(leftAntiNodes)-1].Add(inverseVector)
		if !am.grid.Validate(next) {
			break
		}
		leftAntiNodes = append(leftAntiNodes, next)
	}

	for _, antiNode := range append(rightAntiNodes, leftAntiNodes...) {
		am.antiNodes[antiNode.X][antiNode.Y] = true
	}
}

func (am *AntennaMap) CountAntiNodes() int {
	count := 0
	for _, row := range am.antiNodes {
		count += len(row)
	}
	return count
}
