package rangeset_test

import (
	"math"
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestComplement(t *testing.T) {
	equals := RangeSet.Equals
	assert(t, "Case 1", equals(
		RangeSet{}.Complement(),
		RangeSet{{math.MinInt64, math.MaxInt64}},
	))
	assert(t, "Case 2", equals(
		RangeSet{{math.MinInt64, math.MaxInt64}}.Complement(),
		RangeSet{},
	))
	assert(t, "Case 3", equals(
		RangeSet{{1, 4}, {6, 9}}.Complement(),
		RangeSet{{math.MinInt64, 1}, {4, 6}, {9, math.MaxInt64}},
	))
}
