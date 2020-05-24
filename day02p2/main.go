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

type Computer struct {
	program []int
	memory  []int
}

func (c *Computer) RunProgram(noun, verb int) (int, error) {
	// reset memory
	c.memory = c.memory[:0]
	for _, n := range c.program {
		c.memory = append(c.memory, n)
	}
	// execute
	c.memory[1], c.memory[2] = noun, verb
	if err := execute(c.memory); err != nil {
		return 0, err
	}
	return c.memory[0], nil
}

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

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	prog, err := parseProgram(string(b))
	if err != nil {
		log.Fatal("parsing failed:", err)
	}

	comp := Computer{
		program: prog,
	}

	found := false
	for n := 0; !found && n < 100; n++ {
		for v := 0; !found && v < 100; v++ {
			out, err := comp.RunProgram(n, v)
			if err != nil {
				log.Fatal(err)
			}
			if out == 19690720 {
				found = true
				fmt.Println(n, v)
			}
		}
	}
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
