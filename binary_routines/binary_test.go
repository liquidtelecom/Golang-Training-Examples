// Binary level functions test routines
package main

import (
	"fmt"
	"testing"
)

func TestBitwiseAnd(t *testing.T) {
	tp := struct {
		a, b uint16
		want uint16
	}{5, 6, 4}
	t.Run(fmt.Sprintf("%d,%d", tp.a, tp.b), func(*testing.T) {
		ans := BitwiseAnd(tp.a, tp.b)
		if ans != tp.want {
			t.Errorf("%d [%016b] != %d [%016b]", ans, ans, tp.want, tp.want)
		}
	})
}

func TestBitwiseOr(t *testing.T) {
	tp := struct {
		a, b uint16
		want uint16
	}{5, 6, 7}
	t.Run(fmt.Sprintf("%d,%d", tp.a, tp.b), func(*testing.T) {
		ans := BitwiseOr(tp.a, tp.b)
		if ans != tp.want {
			t.Errorf("%d [%016b] != %d [%016b]", ans, ans, tp.want, tp.want)
		}
	})
}

func TestExclusiveOr(t *testing.T) {
	tp := struct {
		a, b uint16
		want uint16
	}{5, 6, 3}
	t.Run(fmt.Sprintf("%d,%d", tp.a, tp.b), func(*testing.T) {
		ans := ExclusiveOr(tp.a, tp.b)
		if ans != tp.want {
			t.Errorf("%d [%016b] != %d [%016b]", ans, ans, tp.want, tp.want)
		}
	})
}

func TestLeftShift(t *testing.T) {
	tp := struct {
		a, b uint16
		want uint16
	}{8, 2, 32}
	t.Run(fmt.Sprintf("%d,%d", tp.a, tp.b), func(*testing.T) {
		ans := LeftShift(tp.a, tp.b)
		if ans != tp.want {
			t.Errorf("%d [%016b] != %d [%016b]", ans, ans, tp.want, tp.want)
		}
	})
}

func TestRightShift(t *testing.T) {
	tp := struct {
		a, b uint16
		want uint16
	}{8, 2, 2}
	t.Run(fmt.Sprintf("%d,%d", tp.a, tp.b), func(*testing.T) {
		ans := RightShift(tp.a, tp.b)
		if ans != tp.want {
			t.Errorf("%d [%016b] != %d [%016b]", ans, ans, tp.want, tp.want)
		}
	})
}

func TestTestBit(t *testing.T) {
	testName8 := fmt.Sprintf("%d,%d", 193, 0)
	testName16 := fmt.Sprintf("%d,%d", 0xFFFE, 1)
	testName32 := fmt.Sprintf("%d,%d", 0xFFFFFFFE, 6)
	testName64 := fmt.Sprintf("%d,%d", 1, 64)
	testNameErr := fmt.Sprintf("%.2f,1", 1.0)
	t.Run(testName8, func(*testing.T) {
		ans, err := TestBit(uint8(193), 0)
		if ans != true && err != nil {
			t.Errorf("Expected true with nil error, got %v [%v]", ans, err)
		}
	})
	t.Run(testName16, func(*testing.T) {
		ans, err := TestBit(uint16(0xFFFE), 1)
		if ans != true && err != nil {
			t.Errorf("Expected true with nil error, got %v [%v]", ans, err)
		}
	})
	t.Run(testName32, func(*testing.T) {
		ans, err := TestBit(uint32(0xFFFFFFFE), 6)
		if ans != true && err != nil {
			t.Errorf("Expected true with nil error, got %v [%v]", ans, err)
		}
	})
	t.Run(testName64, func(*testing.T) {
		ans, err := TestBit(uint64(1), 0)
		if ans != true && err != nil {
			t.Errorf("Expected true with nil error, got %v [%v]", ans, err)
		}
	})
	t.Run(testNameErr, func(*testing.T) {
		ans, err := TestBit(1.0, 0)
		if ans != false || err == nil {
			t.Errorf("Expected false with error, got %v [%v]", ans, err)
		}
	})
}

func TestMergeInt8Numbers(t *testing.T) {
	type args struct {
		num1 uint8
		num2 uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				num1: 10,
				num2: 10,
			},
			want: "0000101000001010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeInt8Numbers(tt.args.num1, tt.args.num2); fmt.Sprintf("%016b", got) != tt.want {
				t.Errorf("MergeInt8Numbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
