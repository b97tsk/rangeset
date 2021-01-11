package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			RangeSet{{1, 3}}.Intersection(RangeSet{{5, 7}}),
			RangeSet{},
		},
		{
			RangeSet{{1, 5}}.Intersection(RangeSet{{3, 7}}),
			RangeSet{{3, 5}},
		},
		{
			Intersection(
				RangeSet{{3, 11}, {13, 21}},
				RangeSet{{1, 5}, {9, 15}, {19, 23}},
			),
			RangeSet{{3, 5}, {9, 11}, {13, 15}, {19, 21}},
		},
		{
			Intersection(
				RangeSet{{3, 11}, {13, 21}},
				RangeSet{{1, 5}, {9, 15}, {19, 23}},
				RangeSet{{5, 19}},
			),
			RangeSet{{9, 11}, {13, 15}},
		},
		{Intersection(), RangeSet{}},
		{Intersection(RangeSet{}), RangeSet{}},
	}
	for i, c := range testCases {
		if !c.Result.Equals(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}
