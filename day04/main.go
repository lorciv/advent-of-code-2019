package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
)

func check(passwd string, min, max int) error {
	if l := len(passwd); l != 6 {
		return fmt.Errorf("invalid length %d", l)
	}

	value, err := strconv.Atoi(passwd)
	if err != nil {
		return errors.New("not a number")
	}
	if value < min || value > max {
		return errors.New("out of range")
	}

	double := false
	for i := 0; i < len(passwd)-1; i++ {
		if passwd[i] == passwd[i+1] {
			double = true
			break
		}
	}
	if !double {
		return errors.New("no double digit")
	}

	for i := 0; i < len(passwd)-1; i++ {
		cur := passwd[i]
		next := passwd[i+1]
		if next < cur {
			return fmt.Errorf("decreasing pair %c%c", cur, next)
		}
	}

	return nil
}

func check2(passwd string, min, max int) error {
	if err := check(passwd, min, max); err != nil {
		return err
	}

	cur := passwd[0]
	count := 1
	for i := 1; i < len(passwd); i++ {
		if passwd[i] == cur {
			count++
			continue
		}
		if count == 2 {
			break
		}
		cur = passwd[i]
		count = 1
	}

	if count != 2 {
		return errors.New("no group of exactly 2 equal digits")
	}

	return nil
}

var (
	min = flag.Int("min", 165432, "The lowest value in the range for search")
	max = flag.Int("max", 707912, "The highest value in the range for search")
)

func main() {
	flag.Parse()

	count := 0
	count2 := 0
	for i := *min; i <= *max; i++ {
		attempt := strconv.Itoa(i)
		if check(attempt, *min, *max) == nil {
			count++
		}
		if check2(attempt, *min, *max) == nil {
			count2++
		}
	}

	fmt.Printf("count  = %d\n", count)
	fmt.Printf("count2 = %d\n", count2)
}
