package calculator

import (
	"fmt"
	"testing"
)

type testCase struct {
	arg1 int
	arg2 int
	want int
}

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestMultiply(t *testing.T) {
	cases := []testCase{
		{2, 3, 6},
		{10, 5, 50},
		{-8, -3, 24},
		{0, 9, 0},
		{-7, 6, -42},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d*%d=%d", tc.arg1, tc.arg2, tc.want), func(t *testing.T) {
			got := Multiply(tc.arg1, tc.arg2)
			if tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}
