package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestUnion(t *testing.T) {
	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			RangeSet{{1, 3}, {5, 7}}.Union(RangeSet{}),
			RangeSet{{1, 3}, {5, 7}},
		},
		{
			RangeSet{}.Union(RangeSet{{1, 3}, {5, 7}}),
			RangeSet{{1, 3}, {5, 7}},
		},
		{
			RangeSet{{1, 3}}.Union(RangeSet{{5, 7}}),
			RangeSet{{1, 3}, {5, 7}},
		},
		{
			RangeSet{{1, 5}}.Union(RangeSet{{3, 7}}),
			RangeSet{{1, 7}},
		},
		{
			Union(
				RangeSet{{3, 11}, {13, 21}},
				RangeSet{{1, 5}, {9, 15}, {19, 23}},
			),
			RangeSet{{1, 23}},
		},
		{
			Union(
				RangeSet{{3, 11}, {13, 21}},
				RangeSet{{1, 5}, {9, 15}, {19, 23}},
				RangeSet{{5, 19}},
			),
			RangeSet{{1, 23}},
		},
		{Union(), RangeSet{}},
		{Union(RangeSet{}), RangeSet{}},
	}
	for i, c := range testCases {
		if !c.Result.Equals(c.Expect) {
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}
