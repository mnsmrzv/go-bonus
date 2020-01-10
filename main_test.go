package main

import "testing"

func Test_bonus(t *testing.T) {
	tests := []struct {
		name string
		sales []int
		want int
	}{
		{"Have bonus", []int{12_000, 8_000, 15_000, 8_000}, 350},
		{"No bonus", []int{1_000, 100, 2, 8_000}, 0},
	}
	for _, test := range tests {
		got := bonus(test.sales)
		if got != test.want {
			t.Error("For amount of bonus test:", test.name, "got:", got, "want:", test.want)
		}
	}
}