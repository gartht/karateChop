package karatechop

import "testing"

func TestChop(t *testing.T) {
	cases := []struct {
		expected, searchVal int
		input               []int
	}{
		{-1, 3, []int{}},
		{0, 1, []int{1}},
		{-1, 3, []int{1}},
		{0, 1, []int{1}},
		{0, 1, []int{1, 3, 5}},
		{1, 3, []int{1, 3, 5}},
		{2, 5, []int{1, 3, 5}},
		{-1, 0, []int{1, 3, 5}},
		{-1, 2, []int{1, 3, 5}},
		{-1, 4, []int{1, 3, 5}},
		{-1, 6, []int{1, 3, 5}},
		{0, 1, []int{1, 3, 5, 7}},
		{1, 3, []int{1, 3, 5, 7}},
		{2, 5, []int{1, 3, 5, 7}},
		{3, 7, []int{1, 3, 5, 7}},
		{-1, 0, []int{1, 3, 5, 7}},
		{-1, 2, []int{1, 3, 5, 7}},
		{-1, 4, []int{1, 3, 5, 7}},
		{-1, 6, []int{1, 3, 5, 7}},
		{-1, 8, []int{1, 3, 5, 7}},
		{-1, 14, []int{1, 2, 3, 5, 8, 13}},
	}

	for _, c := range cases {
		output := Chop(c.searchVal, c.input)
		if c.expected != output {
			t.Errorf("chop(%d, %d) == %d, want %d", c.searchVal, c.input, output, c.expected)
		}
	}
}
