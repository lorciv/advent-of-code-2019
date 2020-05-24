package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	add  = 1
	mul  = 2
	halt = 99
)

func execute(prog []int) error {
	i := 0
	for prog[i] != halt {
		switch prog[i] {
		case add:
			a, b, c := prog[i+1], prog[i+2], prog[i+3]
			prog[c] = prog[a] + prog[b]
		case mul:
			a, b, c := prog[i+1], prog[i+2], prog[i+3]
			prog[c] = prog[a] * prog[b]
		default:
			return fmt.Errorf("unknown opcode %q", prog[i])
		}
		i += 4
	}
	return nil
}

func restore(prog []int) {
	prog[1] = 12
	prog[2] = 2
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	prog, err := parseProgram(string(b))
	if err != nil {
		log.Fatal("parsing failed:", err)
	}

	restore(prog)
	execute(prog)

	fmt.Println(prog)
}

func parseProgram(text string) ([]int, error) {
	var prog []int
	text = strings.Trim(text, "\n")
	for _, s := range strings.Split(text, ",") {
		cmd, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		prog = append(prog, cmd)
	}
	return prog, nil
}

func mustParseProgram(text string) []int {
	prog, err := parseProgram(text)
	if err != nil {
		panic(err)
	}
	return prog
}
