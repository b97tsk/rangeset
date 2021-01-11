package rangeset_test

import (
	"math"
	"testing"

	. "github.com/b97tsk/rangeset"
)

func TestFromRange(t *testing.T) {
	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			FromRange(1, 5),
			RangeSet{{1, 5}},
		},
		{
			FromRange(5, 1),
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

func TestAdd(t *testing.T) {
	addRange := func(s RangeSet, r Range) RangeSet {
		s.AddRanges(r)
		return s
	}
	addSingle := func(s RangeSet, single int64) RangeSet {
		s.Add(single)
		return s
	}

	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			addRange(RangeSet{{1, 4}, {9, 12}}, Range{5, 8}),
			RangeSet{{1, 4}, {5, 8}, {9, 12}},
		},
		{
			addSingle(RangeSet{{1, 4}, {9, 12}}, 6),
			RangeSet{{1, 4}, {6, 7}, {9, 12}},
		},
		{
			addRange(RangeSet{{1, 4}, {9, 12}}, Range{4, 8}),
			RangeSet{{1, 8}, {9, 12}},
		},
		{
			addRange(RangeSet{{1, 4}, {9, 12}}, Range{5, 9}),
			RangeSet{{1, 4}, {5, 12}},
		},
		{
			addRange(RangeSet{{1, 4}, {9, 12}}, Range{4, 9}),
			RangeSet{{1, 12}},
		},
		{
			addSingle(RangeSet{{1, 4}, {9, 12}}, 10),
			RangeSet{{1, 4}, {9, 12}},
		},
		{
			addRange(RangeSet{{1, 4}, {9, 12}}, Range{9, 12}),
			RangeSet{{1, 4}, {9, 12}},
		},
		{
			addRange(RangeSet{{1, 4}, {9, 12}}, Range{12, 9}),
			RangeSet{{1, 4}, {9, 12}},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equals(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}

func TestDelete(t *testing.T) {
	deleteRange := func(s RangeSet, r Range) RangeSet {
		s.DeleteRanges(r)
		return s
	}
	deleteSingle := func(s RangeSet, single int64) RangeSet {
		s.Delete(single)
		return s
	}

	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{7, 10}),
			RangeSet{{1, 4}, {13, 16}},
		},
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{7, 9}),
			RangeSet{{1, 4}, {9, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{8, 10}),
			RangeSet{{1, 4}, {7, 8}, {13, 16}},
		},
		{
			deleteSingle(RangeSet{{1, 4}, {7, 10}, {13, 16}}, 8),
			RangeSet{{1, 4}, {7, 8}, {9, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{1, 16}),
			RangeSet{},
		},
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{1, 15}),
			RangeSet{{15, 16}},
		},
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{2, 16}),
			RangeSet{{1, 2}},
		},
		{
			deleteSingle(RangeSet{{1, 4}, {7, 10}, {13, 16}}, 5),
			RangeSet{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{4, 7}),
			RangeSet{{1, 4}, {7, 10}, {13, 16}},
		},
		{
			deleteRange(RangeSet{{1, 4}, {7, 10}, {13, 16}}, Range{7, 4}),
			RangeSet{{1, 4}, {7, 10}, {13, 16}},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equals(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}

func TestContains(t *testing.T) {
	s := RangeSet{{1, 3}, {5, 7}}

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
		s.ContainsAny(1, 3) == true,
		s.ContainsAny(3, 5) == false,
		s.ContainsAny(5, 7) == true,
		s.ContainsAny(1, 7) == true,
		s.ContainsAny(1, 1) == false,
		s.ContainsAny(2, 2) == false,
	}
	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestDifference(t *testing.T) {
	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			RangeSet{{1, 5}, {7, 11}}.Difference(RangeSet{{3, 9}}),
			RangeSet{{1, 3}, {9, 11}},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equals(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}

func TestEquals(t *testing.T) {
	testCases := []struct {
		Result, Expect RangeSet
	}{
		{
			RangeSet{{1, 3}, {5, 7}},
			RangeSet{{1, 3}, {5, 7}},
		},
		{
			RangeSet{{1, 3}, {5, 7}},
			RangeSet{{1, 3}, {5, 9}},
		},
		{
			RangeSet{{1, 3}, {5, 7}},
			RangeSet{{1, 3}},
		},
	}
	for i, c := range testCases {
		if !c.Result.Equals(c.Expect) {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}

func TestExtent(t *testing.T) {
	testCases := []struct {
		Result, Expect Range
	}{
		{
			RangeSet{{1, 3}, {5, 7}}.Extent(),
			Range{1, 7},
		},
		{
			Universal().Extent(),
			Range{math.MinInt64, math.MaxInt64},
		},
		{
			RangeSet{}.Extent(),
			Range{},
		},
	}
	for i, c := range testCases {
		if c.Result != c.Expect {
			t.Fail()
			t.Logf("Case %v: want %v, but got %v", i, c.Expect, c.Result)
		}
	}
}

func TestIsSubsetOf(t *testing.T) {
	assertions := []bool{
		RangeSet{}.IsSubsetOf(RangeSet{}) == true,
		RangeSet{{3, 9}}.IsSubsetOf(RangeSet{{1, 11}}) == true,
		RangeSet{{3, 9}}.IsSubsetOf(RangeSet{{1, 5}, {7, 11}}) == false,
	}
	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}

func TestLength(t *testing.T) {
	assertions := []bool{
		RangeSet{}.Length() == 0,
		RangeSet{{1, 4}}.Length() == 3,
		RangeSet{{1, 3}, {5, 7}}.Length() == 4,
	}
	for i, ok := range assertions {
		if !ok {
			t.Fail()
			t.Logf("Case %v: FAILED", i)
		}
	}
}
