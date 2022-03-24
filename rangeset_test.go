package rangeset_test

import (
	"math"
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestFromRange(t *testing.T) {
	type E int

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			FromRange[E](1, 5),
			RangeSet[E]{{1, 5}},
		},
		{
			FromRange[E](5, 1),
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

func TestAdd(t *testing.T) {
	type E int

	addRange := func(s RangeSet[E], r Range[E]) RangeSet[E] {
		s.AddRange(r.Low, r.High)
		return s
	}
	addSingle := func(s RangeSet[E], e E) RangeSet[E] {
		s.Add(e)
		return s
	}

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			addRange(RangeSet[E]{{1, 4}, {9, 12}}, Range[E]{5, 8}),
			RangeSet[E]{{1, 4}, {5, 8}, {9, 12}},
		},
		{
			addSingle(RangeSet[E]{{1, 4}, {9, 12}}, 6),
			RangeSet[E]{{1, 4}, {6, 7}, {9, 12}},
		},
		{
			addRange(RangeSet[E]{{1, 4}, {9, 12}}, Range[E]{4, 8}),
			RangeSet[E]{{1, 8}, {9, 12}},
		},
		{
			addRange(RangeSet[E]{{1, 4}, {9, 12}}, Range[E]{5, 9}),
			RangeSet[E]{{1, 4}, {5, 12}},
		},
		{
			addRange(RangeSet[E]{{1, 4}, {9, 12}}, Range[E]{4, 9}),
			RangeSet[E]{{1, 12}},
		},
		{
			addSingle(RangeSet[E]{{1, 4}, {9, 12}}, 10),
			RangeSet[E]{{1, 4}, {9, 12}},
		},
		{
			addRange(RangeSet[E]{{1, 4}, {9, 12}}, Range[E]{9, 12}),
			RangeSet[E]{{1, 4}, {9, 12}},
		},
		{
			addRange(RangeSet[E]{{1, 4}, {9, 12}}, Range[E]{12, 9}),
			RangeSet[E]{{1, 4}, {9, 12}},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}

func TestDelete(t *testing.T) {
	type E int

	deleteRange := func(s RangeSet[E], r Range[E]) RangeSet[E] {
		s.DeleteRange(r.Low, r.High)
		return s
	}
	deleteSingle := func(s RangeSet[E], e E) RangeSet[E] {
		s.Delete(e)
		return s
	}

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{7, 10}),
			RangeSet[E]{{1, 4}, {13, 16}},
		},
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{7, 9}),
			RangeSet[E]{{1, 4}, {9, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{8, 10}),
			RangeSet[E]{{1, 4}, {7, 8}, {13, 16}},
		},
		{
			deleteSingle(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, 8),
			RangeSet[E]{{1, 4}, {7, 8}, {9, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{1, 16}),
			RangeSet[E]{},
		},
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{1, 15}),
			RangeSet[E]{{15, 16}},
		},
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{2, 16}),
			RangeSet[E]{{1, 2}},
		},
		{
			deleteSingle(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, 5),
			RangeSet[E]{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{4, 7}),
			RangeSet[E]{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet[E]{{1, 4}, {7, 10}, {13, 16}}, Range[E]{7, 4}),
			RangeSet[E]{{1, 4}, {7, 10}, {13, 16}},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}

func TestContains(t *testing.T) {
	type E int

	s := RangeSet[E]{{1, 3}, {5, 7}}

	assertions := []bool{
		s.Contains(0) == false,
		s.Contains(1) == true,
		s.Contains(2) == true,
		s.Contains(3) == false,
		s.Contains(4) == false,
		s.Contains(5) == true,
		s.Contains(6) == true,
		s.Contains(7) == false,
		s.ContainsRange(1, 3) == true,
		s.ContainsRange(3, 5) == false,
		s.ContainsRange(5, 7) == true,
		s.ContainsRange(1, 7) == false,
		s.ContainsRange(1, 1) == false,
		s.ContainsRange(2, 2) == false,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestDifference(t *testing.T) {
	type E int

	testCases := []struct {
		Result, Expected RangeSet[E]
	}{
		{
			RangeSet[E]{{1, 5}, {7, 11}}.Difference(RangeSet[E]{{3, 9}}),
			RangeSet[E]{{1, 3}, {9, 11}},
		},
	}

	for i, c := range testCases {
		if !c.Result.Equal(c.Expected) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}

func TestEqual(t *testing.T) {
	type E int

	assertions := []bool{
		RangeSet[E]{{1, 3}, {5, 7}}.Equal(RangeSet[E]{{1, 3}, {5, 7}}),
		!RangeSet[E]{{1, 3}, {5, 7}}.Equal(RangeSet[E]{{1, 3}, {5, 9}}),
		!RangeSet[E]{{1, 3}, {5, 7}}.Equal(RangeSet[E]{{1, 3}}),
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestExtent(t *testing.T) {
	type E int

	testCases := []struct {
		Result, Expected Range[E]
	}{
		{
			RangeSet[E]{{1, 3}, {5, 7}}.Extent(),
			Range[E]{1, 7},
		},
		{
			Universal[E]().Extent(),
			Range[E]{math.MinInt, math.MaxInt},
		},
		{
			RangeSet[E]{}.Extent(),
			Range[E]{},
		},
	}
	for i, c := range testCases {
		if c.Result != c.Expected {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expected, c.Result)
		}
	}
}

func TestIsSubsetOf(t *testing.T) {
	type E int

	assertions := []bool{
		RangeSet[E]{}.IsSubsetOf(RangeSet[E]{}) == true,
		RangeSet[E]{{3, 9}}.IsSubsetOf(RangeSet[E]{{1, 11}}) == true,
		RangeSet[E]{{3, 9}}.IsSubsetOf(RangeSet[E]{{1, 5}, {7, 11}}) == false,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestCount(t *testing.T) {
	type E int

	assertions := []bool{
		RangeSet[E]{}.Count() == 0,
		RangeSet[E]{{1, 4}}.Count() == 3,
		RangeSet[E]{{1, 3}, {5, 7}}.Count() == 4,
	}

	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}
