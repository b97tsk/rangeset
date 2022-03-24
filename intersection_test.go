package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestIntersection(t *testing.T) {
	type E int

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			RangeSet[E]{{1, 3}}.Intersection(RangeSet[E]{{5, 7}}),
			RangeSet[E]{},
		},
		{
			RangeSet[E]{{1, 5}}.Intersection(RangeSet[E]{{3, 7}}),
			RangeSet[E]{{3, 5}},
		},
		{
			Intersection(
				RangeSet[E]{{3, 11}, {13, 21}},
				RangeSet[E]{{1, 5}, {9, 15}, {19, 23}},
			),
			RangeSet[E]{{3, 5}, {9, 11}, {13, 15}, {19, 21}},
		},
		{
			Intersection(
				RangeSet[E]{{3, 11}, {13, 21}},
				RangeSet[E]{{1, 5}, {9, 15}, {19, 23}},
				RangeSet[E]{{5, 19}},
			),
			RangeSet[E]{{9, 11}, {13, 15}},
		},
		{Intersection[E](), RangeSet[E]{}},
		{Intersection(RangeSet[E]{}), RangeSet[E]{}},
		{
			func() RangeSet[E] {
				var x2, x3, x5 RangeSet[E]

				for i := 2; i < 100; i += 2 {
					x2.Add(E(i))
				}

				for i := 3; i < 100; i += 3 {
					x3.Add(E(i))
				}

				for i := 5; i < 100; i += 5 {
					x5.Add(E(i))
				}

				return Intersection(x2, x3, x5)
			}(),
			RangeSet[E]{{30, 31}, {60, 61}, {90, 91}},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}
