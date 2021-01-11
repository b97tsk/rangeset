package rangeset_test

import (
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestSymmetricDifference(t *testing.T) {
	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			SymmetricDifference(RangeSet{{1, 3}}, RangeSet{{5, 7}}),
			RangeSet{{1, 3}, {5, 7}},
		},
		{
			SymmetricDifference(RangeSet{{1, 5}}, RangeSet{{3, 7}}),
			RangeSet{{1, 3}, {5, 7}},
		},
		{
			SymmetricDifference(RangeSet{{1, 5}}, RangeSet{{5, 9}}),
			RangeSet{{1, 9}},
		},
		{
			SymmetricDifference(
				RangeSet{{3, 11}, {13, 21}},
				RangeSet{{1, 5}, {9, 15}, {19, 23}},
			),
			RangeSet{{1, 3}, {5, 9}, {11, 13}, {15, 19}, {21, 23}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{9, 21}},
			),
			RangeSet{{1, 5}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{9, 19}},
			),
			RangeSet{{1, 5}, {13, 17}, {19, 21}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{9, 23}},
			),
			RangeSet{{1, 5}, {13, 17}, {21, 23}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{9, 25}},
			),
			RangeSet{{1, 5}, {13, 17}, {21, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{5, 21}},
			),
			RangeSet{{1, 9}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{7, 21}},
			),
			RangeSet{{1, 5}, {7, 9}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{11, 21}},
			),
			RangeSet{{1, 5}, {9, 11}, {13, 17}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{7, 23}},
			),
			RangeSet{{1, 5}, {7, 9}, {13, 17}, {21, 23}, {25, 29}},
		},
		{
			SymmetricDifference(
				RangeSet{{1, 5}, {9, 13}, {17, 21}, {25, 29}},
				RangeSet{{5, 25}},
			),
			RangeSet{{1, 9}, {13, 17}, {21, 29}},
		},
		{
			SymmetricDifference(RangeSet{}, RangeSet{}),
			RangeSet{},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equals(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}
