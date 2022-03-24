package rangeset_test

import (
	"math"
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestComplement(t *testing.T) {
	type E int

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			RangeSet[E]{}.Complement(),
			RangeSet[E]{{math.MinInt, math.MaxInt}},
		},
		{
			RangeSet[E]{{math.MinInt, math.MaxInt}}.Complement(),
			RangeSet[E]{},
		},
		{
			RangeSet[E]{{1, 5}, {9, 13}}.Complement(),
			RangeSet[E]{{math.MinInt, 1}, {5, 9}, {13, math.MaxInt}},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}

func TestComplement_unsigned(t *testing.T) {
	type E uint

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			RangeSet[E]{}.Complement(),
			RangeSet[E]{{0, math.MaxUint}},
		},
		{
			RangeSet[E]{{0, math.MaxUint}}.Complement(),
			RangeSet[E]{},
		},
		{
			RangeSet[E]{{1, 5}, {9, 13}}.Complement(),
			RangeSet[E]{{0, 1}, {5, 9}, {13, math.MaxUint}},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}
