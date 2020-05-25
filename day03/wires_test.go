package main

import "testing"

func TestParsePath(t *testing.T) {
	tests := []struct {
		input string
		want  []loc
	}{
		{"D3", []loc{{0, 0}, {0, -1}, {0, -2}, {0, -3}}},
		{"R8,U5,L5,D3", []loc{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0},
			{6, 0}, {7, 0}, {8, 0}, {8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5}, {7, 5},
			{6, 5}, {5, 5}, {4, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}}},
	}

	for _, test := range tests {
		got := mustParsePath(test.input)
		if !samePath(got, test.want) {
			t.Errorf("parsePath(%s) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestCross(t *testing.T) {
	tests := []struct {
		p1, p2 string
		want   []loc
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", []loc{{0, 0}, {3, 3}, {6, 5}}},
	}

	for _, test := range tests {
		p1, p2 := mustParsePath(test.p1), mustParsePath(test.p2)
		got := cross([][]loc{p1, p2})
		if !sameLocs(got, test.want) {
			t.Errorf("cross(%q, %q) = %v, want %v", test.p1, test.p2, got, test.want)
		}
	}
}

func samePath(a, b []loc) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sameLocs(a, b []loc) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(a); j++ {
			if a[i] == b[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func TestSteps(t *testing.T) {
	tests := []struct {
		p      string
		target loc
		want   int
	}{
		{"R8,U5,L5,D3", loc{3, 3}, 20},
		{"R8,U5,L5,D3", loc{6, 5}, 15},
		{"U7,R6,D4,L4", loc{3, 3}, 20},
		{"R8,U5,L5,D3", loc{6, 5}, 15},
	}

	for _, test := range tests {
		path := mustParsePath(test.p)[1:]
		got := steps(path, test.target)
		if got != test.want {
			t.Errorf("steps(%s, %s) = %d, want %d", path, test.target, got, test.want)
		}
	}
}
