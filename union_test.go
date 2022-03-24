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
		{
			func() RangeSet {
				var x2, x3, x5 RangeSet

				for i := 2; i < 30; i += 2 {
					x2.Add(int64(i))
				}

				for i := 3; i < 30; i += 3 {
					x3.Add(int64(i))
				}

				for i := 5; i < 30; i += 5 {
					x5.Add(int64(i))
				}

				return Union(x2, x3, x5)
			}(),
			RangeSet{{2, 7}, {8, 11}, {12, 13}, {14, 17}, {18, 19}, {20, 23}, {24, 29}},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equal(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}
