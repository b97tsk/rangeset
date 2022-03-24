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
		{
			func() RangeSet {
				var x2, x3, x5 RangeSet

				for i := 2; i < 100; i += 2 {
					x2.Add(int64(i))
				}

				for i := 3; i < 100; i += 3 {
					x3.Add(int64(i))
				}

				for i := 5; i < 100; i += 5 {
					x5.Add(int64(i))
				}

				return Intersection(x2, x3, x5)
			}(),
			RangeSet{{30, 31}, {60, 61}, {90, 91}},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equal(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}
