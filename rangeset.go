// Package rangeset is a library for manipulating sets of ranges.
package rangeset

import (
	"math"
	"sort"
)

// A Range is a half-open interval of int64.
type Range struct {
	Low  int64
	High int64 // exclusive
}

// A RangeSet is a non-overlapped ordered slice of Range.
// The zero value for a RangeSet is an empty set ready to use.
//
// Since Range is half-open, you can never add math.MaxInt64 into a RangeSet.
// Thus, complement of an empty set is [math.MinInt64, math.MaxInt64).
type RangeSet []Range

// Add adds a single int64 into set.
func (set *RangeSet) Add(single int64) {
	set.AddRange(single, single+1)
}

// AddRange adds a half-open range into set.
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

	if i == j { // Case 3.
		if low < high {
			s = append(s, Range{})
			copy(s[i+1:], s[i:])
			s[i] = Range{low, high}
			*set = s
		}

		return
	}

	// Case 4 and 5.

	s[i] = Range{low, high}
	s = append(s[:i+1], s[j:]...)
	*set = s
}

// Delete removes a single int64 from set.
func (set *RangeSet) Delete(single int64) {
	set.DeleteRange(single, single+1)
}

// DeleteRange removes a half-open range from set.
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

// Contains reports whether or not set contains a single int64.
func (set RangeSet) Contains(single int64) bool {
	return set.ContainsRange(single, single+1)
}

// ContainsRange reports whether or not set contains every int64 in a half-
// open range.
func (set RangeSet) ContainsRange(low, high int64) bool {
	i := sort.Search(len(set), func(i int) bool { return set[i].High > low })
	return i < len(set) && set[i].Low <= low && high <= set[i].High && low < high
}

// ContainsAny reports whether or not set contains any int64 in a half-open
// range.
func (set RangeSet) ContainsAny(low, high int64) bool {
	i := sort.Search(len(set), func(i int) bool { return set[i].High > low })
	t := set[i:]
	j := i + sort.Search(len(t), func(i int) bool { return t[i].Low >= high })

	return i < j && low < high
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

// Complement returns the inverse of set.
func (set RangeSet) Complement() RangeSet {
	if len(set) == 0 {
		return RangeSet{{math.MinInt64, math.MaxInt64}}
	}

	return set.complement()
}

func (set RangeSet) complement() RangeSet {
	var result RangeSet

	if len(set) > 1 {
		result = make(RangeSet, 0, len(set)+1) // Pre-allocation.
	}

	r0 := set[0]

	if r0.Low > math.MinInt64 {
		result = append(result, Range{math.MinInt64, r0.Low})
	}

	low := r0.High

	for _, r := range set[1:] {
		result = append(result, Range{low, r.Low})
		low = r.High
	}

	if low < math.MaxInt64 {
		result = append(result, Range{low, math.MaxInt64})
	}

	return result
}

// Length returns the number of int64 in set.
func (set RangeSet) Length() uint64 {
	acc := uint64(0)
	for _, r := range set {
		acc += uint64(r.High - r.Low)
	}

	return acc
}
