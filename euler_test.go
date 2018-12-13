package euler

import (
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	for i := 1; i <= len(problems); i++ {
		t.Run(fmt.Sprintf("Solution %d", i), func(t *testing.T) {
			got := Solve(i)
			expected := GetSolution(i)
			if got != expected {
				t.Errorf("Failed to solve problem %d: Got %s, expected %s", i, got, expected)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 1; i <= len(problems); i++ {
		b.Run(fmt.Sprintf("Solution %d", i), func(b *testing.B) {
			Solve(i)
		})
	}
}
