package rangeset_test

import (
	"math"
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestComplement(t *testing.T) {
	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			RangeSet{}.Complement(),
			RangeSet{{math.MinInt64, math.MaxInt64}},
		},
		{
			RangeSet{{math.MinInt64, math.MaxInt64}}.Complement(),
			RangeSet{},
		},
		{
			RangeSet{{1, 5}, {9, 13}}.Complement(),
			RangeSet{{math.MinInt64, 1}, {5, 9}, {13, math.MaxInt64}},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equal(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}
