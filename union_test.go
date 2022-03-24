package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestUnion(t *testing.T) {
	type E int

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			RangeSet[E]{{1, 3}, {5, 7}}.Union(RangeSet[E]{}),
			RangeSet[E]{{1, 3}, {5, 7}},
		},
		{
			RangeSet[E]{}.Union(RangeSet[E]{{1, 3}, {5, 7}}),
			RangeSet[E]{{1, 3}, {5, 7}},
		},
		{
			RangeSet[E]{{1, 3}}.Union(RangeSet[E]{{5, 7}}),
			RangeSet[E]{{1, 3}, {5, 7}},
		},
		{
			RangeSet[E]{{1, 5}}.Union(RangeSet[E]{{3, 7}}),
			RangeSet[E]{{1, 7}},
		},
		{
			Union(
				RangeSet[E]{{3, 11}, {13, 21}},
				RangeSet[E]{{1, 5}, {9, 15}, {19, 23}},
			),
			RangeSet[E]{{1, 23}},
		},
		{
			Union(
				RangeSet[E]{{3, 11}, {13, 21}},
				RangeSet[E]{{1, 5}, {9, 15}, {19, 23}},
				RangeSet[E]{{5, 19}},
			),
			RangeSet[E]{{1, 23}},
		},
		{Union[E](), RangeSet[E]{}},
		{Union(RangeSet[E]{}), RangeSet[E]{}},
		{
			func() RangeSet[E] {
				var x2, x3, x5 RangeSet[E]

				for i := 2; i < 30; i += 2 {
					x2.Add(E(i))
				}

				for i := 3; i < 30; i += 3 {
					x3.Add(E(i))
				}

				for i := 5; i < 30; i += 5 {
					x5.Add(E(i))
				}

				return Union(x2, x3, x5)
			}(),
			RangeSet[E]{{2, 7}, {8, 11}, {12, 13}, {14, 17}, {18, 19}, {20, 23}, {24, 29}},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}
