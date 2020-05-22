package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fuel(mass int) int {
	return mass/3 - 2
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	totFuel := 0

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("skipping invalid mass %s", line)
			continue
		}
		totFuel += fuel(mass)
	}

	fmt.Println(totFuel)
}
