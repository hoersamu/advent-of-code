package trailfinder

import (
	"bufio"
	"common"
)

const MIN_ELEVATION = 0
const MAX_ELEVATION = 9

var matrix = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type TrailFinder struct {
	topo_map [][]int
	grid     common.Grid
}

func NewTrailFinder(scanner *bufio.Scanner) *TrailFinder {
	topo_map := common.ScanWithDelimitersAsInt(scanner, "")

	return &TrailFinder{
		topo_map: topo_map,
		grid:     common.Grid{SizeX: len(topo_map[0]), SizeY: len(topo_map)},
	}
}

func (tf *TrailFinder) getElevation(position common.Point) int {
	return tf.topo_map[position.Y][position.X]
}

func (tf *TrailFinder) findTrailEnd(position common.Point) (map[string]bool, int) {
	elevation := tf.getElevation(position)
	if elevation == MAX_ELEVATION {
		return map[string]bool{position.String(): true}, 1
	}

	trailEndMap := map[string]bool{}
	trailEndCount := 0

	for _, direction := range matrix {
		nextPosition := common.Point{X: position.X + direction[0], Y: position.Y + direction[1]}

		if tf.grid.Validate(nextPosition) && tf.getElevation(nextPosition) == elevation+1 {
			endMap, count := tf.findTrailEnd(nextPosition)
			trailEndCount += count
			for k := range endMap {
				trailEndMap[k] = true
			}
		}
	}

	return trailEndMap, trailEndCount
}

func (tf *TrailFinder) FindTrails() (int, int) {
	trailEnds := 0
	trailCount := 0

	for y, row := range tf.topo_map {
		for x := range row {
			if tf.getElevation(common.Point{X: x, Y: y}) == MIN_ELEVATION {
				trailEndMap, trailsCount := tf.findTrailEnd(common.Point{X: x, Y: y})
				trailEnds += len(trailEndMap)
				trailCount += trailsCount
			}
		}
	}

	return trailEnds, trailCount
}
