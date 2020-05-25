package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parsePath(text string) ([]loc, error) {
	path := []loc{{0, 0}}
	steps := strings.Split(text, ",")

	for _, step := range steps {
		num, err := strconv.Atoi(step[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid amount %q", step[1:])
		}

		var dir loc
		switch step[:1] {
		case "R":
			dir = loc{1, 0}
		case "U":
			dir = loc{0, 1}
		case "L":
			dir = loc{-1, 0}
		case "D":
			dir = loc{0, -1}
		default:
			return nil, fmt.Errorf("invalid direction %q", step[:1])
		}

		for i := 0; i < num; i++ {
			lastloc := path[len(path)-1]
			nextloc := loc{
				x: lastloc.x + dir.x,
				y: lastloc.y + dir.y,
			}
			path = append(path, nextloc)
		}
	}

	return path, nil
}

func mustParsePath(text string) []loc {
	path, err := parsePath(text)
	if err != nil {
		panic(err)
	}
	return path
}

func cross(paths [][]loc) []loc {
	locs := make(map[loc]int)
	for _, path := range paths {
		uniq := make(map[loc]struct{})
		for _, l := range path {
			uniq[l] = struct{}{}
		}
		for l := range uniq {
			locs[l]++
		}
	}

	var crosses []loc
	for l, count := range locs {
		if count == len(paths) {
			crosses = append(crosses, l)
		}
	}
	return crosses
}

func steps(wire []loc, target loc) int {
	count := 0
	for _, step := range wire {
		count++
		if step == target {
			break
		}
	}
	return count
}

var file = flag.String("file", "input.txt", "file containing the path of each wire, one wire per line of text")

func main() {
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var wires [][]loc

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		w, err := parsePath(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		wires = append(wires, w[1:])
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// part 1
	crosses := cross(wires)
	fmt.Println("distance of closest cross:", loc{0, 0}.minDistance(crosses))

	// part 2
	minSteps := math.MaxInt64

	for _, c := range crosses {
		curSteps := 0
		for _, w := range wires {
			curSteps += steps(w, c)
		}
		if curSteps < minSteps {
			minSteps = curSteps
		}
	}

	fmt.Println("fewest combined steps:", minSteps)
}
