package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestIntersect(t *testing.T) {
	equals := RangeSet.Equals
	assert(t, "Case 1", equals(
		RangeSet{{1, 3}}.Intersect(RangeSet{{5, 7}}),
		RangeSet{},
	))
	assert(t, "Case 2", equals(
		RangeSet{{1, 5}}.Intersect(RangeSet{{3, 7}}),
		RangeSet{{3, 5}},
	))
	assert(t, "Case 3", equals(
		Intersect(
			RangeSet{{2, 6}, {7, 12}},
			RangeSet{{1, 4}, {5, 9}, {10, 16}},
		),
		RangeSet{{2, 4}, {5, 6}, {7, 9}, {10, 12}},
	))
	assert(t, "Case 4", equals(
		Intersect(
			RangeSet{{2, 6}, {7, 12}},
			RangeSet{{1, 4}, {5, 9}, {10, 16}},
			RangeSet{{1, 11}},
		),
		RangeSet{{2, 4}, {5, 6}, {7, 9}, {10, 11}},
	))
	assert(t, "Case 5", equals(Intersect(), RangeSet{}))
	assert(t, "Case 6", equals(Intersect(RangeSet{}), RangeSet{}))
}
