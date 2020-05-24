package main

import "testing"

func TestFuel(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, test := range tests {
		got := fuel(test.mass)
		if got != test.want {
			t.Errorf("fuel(%d) = %d, want %d", test.mass, got, test.want)
		}
	}
}
