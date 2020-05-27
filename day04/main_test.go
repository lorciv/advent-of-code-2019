package main

import "testing"

func TestCheck(t *testing.T) {
	const min = 165432
	const max = 707912

	tests := []struct {
		passwd string
		ok     bool
	}{
		{"18g37rg317g731", false},
		{"18g37r", false},
		{"978989", false},
		{"234543", false},
		{"234567", false},
		{"233456", true},
		{"333333", true},
	}

	for _, test := range tests {
		got := check(test.passwd, min, max)
		if test.ok && got != nil {
			t.Errorf("check(%s) = error: %v, want ok", test.passwd, got)
		} else if !test.ok && got == nil {
			t.Errorf("check(%s) = ok, want error", test.passwd)
		}
	}
}

func TestCheck2(t *testing.T) {
	const min = 165432
	const max = 707912

	tests := []struct {
		passwd string
		ok     bool
	}{
		{"18g37rg317g731", false},
		{"18g37r", false},
		{"978989", false},
		{"234543", false},
		{"234567", false},
		{"233456", true},
		{"333333", false},
		{"223344", true},
		{"123444", false},
		{"222233", true},
	}

	for _, test := range tests {
		got := check2(test.passwd, min, max)
		if test.ok && got != nil {
			t.Errorf("check2(%s) = error: %v, want ok", test.passwd, got)
		} else if !test.ok && got == nil {
			t.Errorf("check2(%s) = ok, want error", test.passwd)
		}
	}
}
