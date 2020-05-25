package main

import (
	"fmt"
	"math"
)

type loc struct {
	x, y int
}

func (l loc) manhattan(other loc) int {
	return abs(l.x-other.x) + abs(l.y-other.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (l loc) minDistance(others []loc) int {
	mindist := math.MaxInt32
	for _, other := range others {
		dist := l.manhattan(other)
		if dist < mindist {
			mindist = dist
		}
	}
	return mindist
}

func (l loc) String() string {
	return fmt.Sprintf("(%d, %d)", l.x, l.y)
}
