// Package rangeset is a library for manipulating sets of non-overlapping
// non-adjacent ranges.
package rangeset

import (
	"math"
	"sort"
)

// A Range is a half-open interval of int64.
type Range struct {
	Low  int64 // inclusive
	High int64 // exclusive
}

// A RangeSet is a slice of non-overlapping non-adjacent ordered Ranges.
// The zero value for a RangeSet, i.e. a nil RangeSet, is an empty set.
//
// Since Range is half-open, you can never add math.MaxInt64 into a RangeSet.
// Thus, complement of an empty set is [math.MinInt64, math.MaxInt64).
type RangeSet []Range

// FromRange creates a RangeSet from range [low, high).
//
// If low >= high, FromRange returns nil.
func FromRange(low, high int64) RangeSet {
	if low >= high {
		return nil
	}

	return RangeSet{{low, high}}
}

// Universal returns the largest RangeSet, which is the complement of an empty
// set, i.e. [math.MinInt64, math.MaxInt64).
func Universal() RangeSet {
	return FromRange(math.MinInt64, math.MaxInt64)
}

// Add adds a single int64 into set.
func (set *RangeSet) Add(single int64) {
	set.AddRange(single, single+1)
}

// AddRange adds range [low, high) into set.
func (set *RangeSet) AddRange(low, high int64) {
	s := *set

	i := sort.Search(len(s), func(i int) bool { return s[i].Low > low })
	j := sort.Search(len(s), func(i int) bool { return s[i].High > high })

	// ┌────────┬─────────────────────────────────────────┐
	// │        │    j-1        j        i-1        i     │
	// │ Case 1 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │        |<-high ->|   |<- low ->|        │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    j-1        j         i               │
	// │ Case 2 │  |-----|   |-----|   |-----|            │
	// │        │            |<- low ->|                  │
	// │        │        |<-high ->|                      │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1       i,j                        │
	// │ Case 3 │  |-----|   |-----|                      │
	// │        │  |<- low ->|                            │
	// │        │        |<-high ->|                      │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i         j               │
	// │ Case 4 │  |-----|   |-----|   |-----|            │
	// │        │  |<- low ->|     |<-high ->|            │
	// │        │                                         │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i        j-1        j     │
	// │ Case 5 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │  |<- low ->|               |<-high ->|  │
	// └────────┴─────────────────────────────────────────┘

	if i > j { // Case 1 and 2.
		return
	}

	// Case 3, 4 and 5.

	if i > 0 && low <= s[i-1].High {
		low = s[i-1].Low
		i--
	}

	if j < len(s) && high >= s[j].Low {
		high = s[j].High
		j++
	}

	if i == j { // Case 3 (where low and high overlaps).
		if low < high {
			s = append(s, Range{})
			copy(s[i+1:], s[i:])
			s[i] = Range{low, high}
			*set = s
		}

		return
	}

	s[i] = Range{low, high}
	s = append(s[:i+1], s[j:]...)
	*set = s
}

// AddRanges adds zero or more half-open ranges into set.
func (set *RangeSet) AddRanges(ranges ...Range) {
	for _, r := range ranges {
		set.AddRange(r.Low, r.High)
	}
}

// Delete removes a single int64 from set.
func (set *RangeSet) Delete(single int64) {
	set.DeleteRange(single, single+1)
}

// DeleteRange removes range [low, high) from set.
func (set *RangeSet) DeleteRange(low, high int64) {
	s := *set

	i := sort.Search(len(s), func(i int) bool { return s[i].High > low })
	// j := sort.Search(len(s), func(i int) bool { return s[i].Low > high })

	// ┌────────┬─────────────────────────────────────────┐
	// │        │    j-1        j        i-1        i     │
	// │ Case 1 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │  |<-high ->|               |<- low ->|  │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    j-1        j         i               │
	// │ Case 2 │  |-----|   |-----|   |-----|            │
	// │        │  |<-high ->|     |<- low ->|            │
	// │        │                                         │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1       i,j                        │
	// │ Case 3 │  |-----|   |-----|                      │
	// │        │        |<- low ->|                      │
	// │        │  |<-high ->|                            │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i         j               │
	// │ Case 4 │  |-----|   |-----|   |-----|            │
	// │        │        |<- low ->|                      │
	// │        │            |<-high ->|                  │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i        j-1        j     │
	// │ Case 5 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │        |<- low ->|   |<-high ->|        │
	// └────────┴─────────────────────────────────────────┘

	// Optimized, j >= i.
	t := s[i:]
	j := i + sort.Search(len(t), func(i int) bool { return t[i].Low > high })

	if i == j { // Case 1, 2 and 3.
		return
	}

	if i == j-1 { // Case 4.
		if low > s[i].Low {
			if high < s[i].High {
				if low < high {
					s = append(s, Range{})
					copy(s[j:], s[i:])
					s[i].High = low
					s[j].Low = high
					*set = s
				}
			} else {
				s[i].High = low
			}
		} else {
			if high < s[i].High {
				s[i].Low = high
			} else {
				s = append(s[:i], s[j:]...)
				*set = s
			}
		}

		return
	}

	// Case 5.

	if low > s[i].Low {
		s[i].High = low
		i++
	}

	if high < s[j-1].High {
		s[j-1].Low = high
		j--
	}

	s = append(s[:i], s[j:]...)
	*set = s
}

// DeleteRanges removes zero or more half-open ranges from set.
func (set *RangeSet) DeleteRanges(ranges ...Range) {
	for _, r := range ranges {
		set.DeleteRange(r.Low, r.High)
	}
}

// Contains reports whether or not set contains a single int64.
func (set RangeSet) Contains(single int64) bool {
	return set.ContainsRange(single, single+1)
}

// ContainsRange reports whether or not set contains every int64 in range
// [low, high).
func (set RangeSet) ContainsRange(low, high int64) bool {
	i := sort.Search(len(set), func(i int) bool { return set[i].High > low })
	return i < len(set) && set[i].Low <= low && high <= set[i].High && low < high
}

// ContainsAny reports whether or not set contains any int64 in range
// [low, high).
func (set RangeSet) ContainsAny(low, high int64) bool {
	i := sort.Search(len(set), func(i int) bool { return set[i].High > low })
	t := set[i:]
	j := i + sort.Search(len(t), func(i int) bool { return t[i].Low >= high })

	return i < j && low < high
}

// Difference returns the subset of set that having all elements in other
// excluded.
func (set RangeSet) Difference(other RangeSet) RangeSet {
	return set.Intersection(other.Complement())
}

// Equals reports whether or not set is identical to other.
func (set RangeSet) Equals(other RangeSet) bool {
	if len(set) != len(other) {
		return false
	}

	for i, r := range set {
		if r != other[i] {
			return false
		}
	}

	return true
}

// Extent returns the smallest Range that covers the whole set.
//
// If set is empty, Extent returns the zero value.
func (set RangeSet) Extent() Range {
	if len(set) == 0 {
		return Range{}
	}

	return Range{
		Low:  set[0].Low,
		High: set[len(set)-1].High,
	}
}

// IsSubsetOf reports whether or not other contains every int64 in set.
func (set RangeSet) IsSubsetOf(other RangeSet) bool {
	for _, r := range set {
		if !other.ContainsRange(r.Low, r.High) {
			return false
		}
	}

	return true
}

// Count returns the number of int64 in set.
func (set RangeSet) Count() uint64 {
	var count uint64

	for _, r := range set {
		count += uint64(r.High - r.Low)
	}

	return count
}
