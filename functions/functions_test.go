// Functions and methods test routines
package main

import (
	"fmt"
	"testing"
)

func TestBasicFunction(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 3, 3},
		{2, 3, 6},
		{3, 3, 9},
		{5, 3, 15},
		{5, 5, 25},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := BasicFunction(tt.a, tt.b)
			if tt.want != ans {
				t.Errorf("got %d while wanting %d", ans, tt.want)
			}
		})
	}
}

func TestBasicFunctionDualReturn(t *testing.T) {
	tests := []struct {
		a, b  int
		want1 bool
		want2 error
	}{
		{1, 1, false, fmt.Errorf("error, a and b were equal")},
		{1, 2, false, nil},
		{2, 1, true, nil},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			b, e := BasicFunctionDualReturn(tt.a, tt.b)
			if b != tt.want1 {
				t.Errorf("wanted boolean %v got %v", b, tt.want1)
			}
			if e != nil && e.Error() != tt.want2.Error() {
				t.Errorf("wanted %v got %v", e, tt.want2)
			}
		})
	}
}

func TestNestedAnonymous(t *testing.T) {
	tests := []struct {
		a, b         int
		want1, want2 int
	}{
		{10, 5, 50, 2},
		{20, 5, 100, 4},
		{100, 10, 1000, 10},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans1, ans2 := NestedAnonymous(tt.a, tt.b)
			if tt.want1 != ans1 || tt.want2 != ans2 {
				t.Errorf("got %d and %d while wanting %d and %d", ans1, ans2, tt.want1, tt.want2)
			}
		})
	}
}

func TestDemoStruct_Multiply(t *testing.T) {
	tests := []struct {
		a    *DemoStruct
		want int
	}{
		{&DemoStruct{a: 1, b: 2}, 2},
		{&DemoStruct{a: 5, b: 10}, 50},
		{&DemoStruct{a: 7, b: 3}, 21},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%+v", tt.a)
		t.Run(testName, func(t *testing.T) {
			ans := tt.a.Multiply()
			if tt.want != ans {
				t.Errorf("got %d while wanting %d", ans, tt.want)
			}
		})
	}
}

func TestSliceReceiver_Append(t *testing.T) {
	tests := []struct {
		sr   *SliceReceiver
		a    []int
		want []int
	}{
		{&SliceReceiver{1, 2}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{&SliceReceiver{1, 2, 3, 4}, []int{5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{&SliceReceiver{1, 3, 5}, []int{7, 9, 11}, []int{1, 3, 5, 7, 9, 11}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v,%v", *tt.sr, tt.a)
		t.Run(testName, func(t *testing.T) {
			tt.sr.Append(tt.a...)
			ans := []int(*tt.sr)
			if len(ans) != len(tt.want) {
				t.Errorf("length of receiver != %d", len(tt.want))
			}
			for a := range ans {
				if ans[a] != tt.want[a] {
					t.Errorf("%d at offset %d != %d", ans[a], a, tt.want[a])
				}
			}
		})
	}
}

func TestSliceReceiver_DeleteElementNoOrder(t *testing.T) {
	tests := []struct {
		sr   *SliceReceiver
		a    int
		want []int
	}{
		{&SliceReceiver{1, 5, 7, 9, 11}, 1, []int{1, 11, 7, 9}},
		{&SliceReceiver{1, 5, 7, 9, 11}, 4, []int{1, 5, 7, 9}},
		{&SliceReceiver{1, 5, 7, 9, 11}, 0, []int{11, 5, 7, 9}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v,%d", *tt.sr, tt.a)
		t.Run(testName, func(*testing.T) {
			err := tt.sr.DeleteElementNoOrder(tt.a)
			if err != nil {
				t.Errorf("error: %v", err)
				return
			}
			ans := []int(*tt.sr)
			if len(ans) != len(tt.want) {
				t.Errorf("length of receiver != %d", len(tt.want))
				return
			}
			for i := 0; i < len(ans)-1; i++ {
				if ans[i] != tt.want[i] {
					t.Errorf("%d at offset %d != %d", ans[i], i, tt.want[i])
					return
				}
			}
		})
	}
}
