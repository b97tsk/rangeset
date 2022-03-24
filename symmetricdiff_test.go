package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestSymmetricDifference(t *testing.T) {
	type E int

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			SymmetricDifference(RangeSet[E]{{1, 3}}, RangeSet[E]{{5, 7}}),
			RangeSet[E]{{1, 3}, {5, 7}},
		},
		{
			SymmetricDifference(RangeSet[E]{{1, 5}}, RangeSet[E]{{3, 7}}),
			RangeSet[E]{{1, 3}, {5, 7}},
		},
		{
			SymmetricDifference(RangeSet[E]{{1, 5}}, RangeSet[E]{{5, 9}}),
			RangeSet[E]{{1, 9}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{3, 11}, {13, 21}},
				RangeSet[E]{{1, 5}, {9, 15}, {19, 23}},
			),
			RangeSet[E]{{1, 3}, {5, 9}, {11, 13}, {15, 19}, {21, 23}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{9, 21}},
			),
			RangeSet[E]{{1, 5}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{9, 19}},
			),
			RangeSet[E]{{1, 5}, {13, 17}, {19, 21}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{9, 23}},
			),
			RangeSet[E]{{1, 5}, {13, 17}, {21, 23}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{9, 25}},
			),
			RangeSet[E]{{1, 5}, {13, 17}, {21, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{5, 21}},
			),
			RangeSet[E]{{1, 9}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{7, 21}},
			),
			RangeSet[E]{{1, 5}, {7, 9}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{11, 21}},
			),
			RangeSet[E]{{1, 5}, {9, 11}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{7, 23}},
			),
			RangeSet[E]{{1, 5}, {7, 9}, {13, 17}, {21, 23}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet[E]{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet[E]{{5, 25}},
			),
			RangeSet[E]{{1, 9}, {13, 17}, {21, 29}},
		},
		{
			SymmetricDifference(RangeSet[E]{}, RangeSet[E]{}),
			RangeSet[E]{},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}
