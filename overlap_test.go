package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestOverlaps(t *testing.T) {
	assertions := []bool{
		RangeSet{}.Overlaps(RangeSet{}) == false,
		RangeSet{{1, 3}}.Overlaps(RangeSet{{5, 7}}) == false,
		RangeSet{{1, 5}}.Overlaps(RangeSet{{3, 7}}) == true,
		RangeSet{{5, 7}}.Overlaps(RangeSet{{1, 3}, {9, 11}}) == false,
		RangeSet{{3, 9}}.Overlaps(RangeSet{{1, 5}, {7, 11}}) == true,
	}
	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}
