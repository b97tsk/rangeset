package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestOverlaps(t *testing.T) {
	type E int

	assertions := []bool{
		RangeSet[E]{}.Overlaps(RangeSet[E]{}) == false,
		RangeSet[E]{{1, 3}}.Overlaps(RangeSet[E]{{5, 7}}) == false,
		RangeSet[E]{{1, 5}}.Overlaps(RangeSet[E]{{3, 7}}) == true,
		RangeSet[E]{{5, 7}}.Overlaps(RangeSet[E]{{1, 3}, {9, 11}}) == false,
		RangeSet[E]{{3, 9}}.Overlaps(RangeSet[E]{{1, 5}, {7, 11}}) == true,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}
