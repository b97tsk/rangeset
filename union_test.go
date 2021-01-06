package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestUnion(t *testing.T) {
	equals := RangeSet.Equals
	assert(t, "Case 1", equals(
		RangeSet{{1, 3}, {5, 7}}.Union(RangeSet{}),
		RangeSet{{1, 3}, {5, 7}},
	))
	assert(t, "Case 2", equals(
		RangeSet{}.Union(RangeSet{{1, 3}, {5, 7}}),
		RangeSet{{1, 3}, {5, 7}},
	))
	assert(t, "Case 3", equals(
		RangeSet{{1, 3}}.Union(RangeSet{{5, 7}}),
		RangeSet{{1, 3}, {5, 7}},
	))
	assert(t, "Case 4", equals(
		RangeSet{{1, 5}}.Union(RangeSet{{3, 7}}),
		RangeSet{{1, 7}},
	))
	assert(t, "Case 5", equals(
		Union(
			RangeSet{{2, 6}, {7, 12}},
			RangeSet{{1, 4}, {5, 9}, {10, 16}},
		),
		RangeSet{{1, 16}},
	))
	assert(t, "Case 6", equals(
		Union(
			RangeSet{{2, 6}, {7, 12}},
			RangeSet{{1, 4}, {5, 9}, {10, 16}},
			RangeSet{{1, 20}},
		),
		RangeSet{{1, 20}},
	))
	assert(t, "Case 7", equals(Union(), RangeSet{}))
}
