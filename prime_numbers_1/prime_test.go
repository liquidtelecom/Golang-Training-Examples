// Prime number test routine
package main

import (
	"fmt"
	"testing"
)

func TestFindPrime(t *testing.T) {
	var tests = []struct {
		a, b int
		want []int
	}{
		{2, 20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{20, 40, []int{23, 29, 31, 17}},
		{40, 60, []int{41, 43, 47, 53, 59}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Min: %2d\tMax: %2d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := FindPrime(tt.a, tt.b)
			if len(ans) != len(tt.want) {
				t.Errorf("got %v instead of %v", ans, tt.want)
			}
			for i := 0; i < len(ans)-1; i++ {
				if ans[i] != tt.want[i] {
					t.Errorf("got %v instead of %v", ans, tt.want)
				}
			}
		})
	}
}
