package main

import "testing"

func TestFuel(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tests {
		got := fuel(test.mass)
		if got != test.want {
			t.Errorf("fuel(%d) = %d, want %d", test.mass, got, test.want)
		}
	}
}
