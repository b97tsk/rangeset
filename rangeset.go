// Package rangeset is a library for manipulating sets of discrete ranges.
package rangeset

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Elem is the type set containing all supported element types.
type Elem constraints.Integer

// A Range is a half-open interval of type E.
type Range[E Elem] struct {
	Low  E // inclusive
	High E // exclusive
}

// A RangeSet is a slice of discrete Ranges sorted in ascending order.
// The zero value for a RangeSet, i.e. a nil RangeSet, is an empty set.
//
// Since Range is half-open, you can never add the maximum value of E into
// a RangeSet.
type RangeSet[E Elem] []Range[E]

// FromRange creates a RangeSet from range [lo, hi).
//
// If lo >= hi, FromRange returns nil.
func FromRange[E Elem](lo, hi E) RangeSet[E] {
	if lo >= hi {
		return nil
	}

	return RangeSet[E]{{lo, hi}}
}

// Universal returns the largest RangeSet, which contains every E except one,
// the maximum value of E.
func Universal[E Elem]() RangeSet[E] {
	return FromRange(minOf[E](), maxOf[E]())
}

// Add adds a single element into set.
func (set *RangeSet[E]) Add(e E) {
	set.AddRange(e, e+1)
}

// AddRange adds range [lo, hi) into set.
func (set *RangeSet[E]) AddRange(lo, hi E) {
	s := *set

	i := sort.Search(len(s), func(i int) bool { return s[i].Low > lo })
	j := sort.Search(len(s), func(i int) bool { return s[i].High > hi })

	// ┌────────┬─────────────────────────────────────────┐
	// │        │    j-1        j        i-1        i     │
	// │ Case 1 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │        |<- hi  ->|   |<- lo  ->|        │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    j-1        j         i               │
	// │ Case 2 │  |-----|   |-----|   |-----|            │
	// │        │            |<- lo  ->|                  │
	// │        │        |<- hi  ->|                      │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1       i,j                        │
	// │ Case 3 │  |-----|   |-----|                      │
	// │        │  |<- lo  ->|                            │
	// │        │        |<- hi  ->|                      │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i         j               │
	// │ Case 4 │  |-----|   |-----|   |-----|            │
	// │        │  |<- lo  ->|     |<- hi  ->|            │
	// │        │                                         │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i        j-1        j     │
	// │ Case 5 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │  |<- lo  ->|               |<- hi  ->|  │
	// └────────┴─────────────────────────────────────────┘

	if i > j { // Case 1 and 2.
		return
	}

	// Case 3, 4 and 5.

	if i > 0 && lo <= s[i-1].High {
		lo = s[i-1].Low
		i--
	}

	if j < len(s) && hi >= s[j].Low {
		hi = s[j].High
		j++
	}

	if i == j { // Case 3 (where lo and hi overlaps).
		if lo < hi {
			s = append(s, Range[E]{})
			copy(s[i+1:], s[i:])
			s[i] = Range[E]{lo, hi}
			*set = s
		}

		return
	}

	s[i] = Range[E]{lo, hi}
	s = append(s[:i+1], s[j:]...)
	*set = s
}

// Delete removes a single element from set.
func (set *RangeSet[E]) Delete(e E) {
	set.DeleteRange(e, e+1)
}

// DeleteRange removes range [lo, hi) from set.
func (set *RangeSet[E]) DeleteRange(lo, hi E) {
	s := *set

	i := sort.Search(len(s), func(i int) bool { return s[i].High > lo })
	// j := sort.Search(len(s), func(i int) bool { return s[i].Low > hi })

	// ┌────────┬─────────────────────────────────────────┐
	// │        │    j-1        j        i-1        i     │
	// │ Case 1 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │  |<- hi  ->|               |<- lo  ->|  │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    j-1        j         i               │
	// │ Case 2 │  |-----|   |-----|   |-----|            │
	// │        │  |<- hi  ->|     |<- lo  ->|            │
	// │        │                                         │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1       i,j                        │
	// │ Case 3 │  |-----|   |-----|                      │
	// │        │        |<- lo  ->|                      │
	// │        │  |<- hi  ->|                            │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i         j               │
	// │ Case 4 │  |-----|   |-----|   |-----|            │
	// │        │        |<- lo  ->|                      │
	// │        │            |<- hi  ->|                  │
	// ├────────┼─────────────────────────────────────────┤
	// │        │    i-1        i        j-1        j     │
	// │ Case 5 │  |-----|   |-----| ~ |-----|   |-----|  │
	// │        │        |<- lo  ->|   |<- hi  ->|        │
	// └────────┴─────────────────────────────────────────┘

	// Optimized, j >= i.
	t := s[i:]
	j := i + sort.Search(len(t), func(i int) bool { return t[i].Low > hi })

	if i == j { // Case 1, 2 and 3.
		return
	}

	if i == j-1 { // Case 4.
		if lo > s[i].Low {
			if hi < s[i].High {
				if lo < hi {
					s = append(s, Range[E]{})
					copy(s[j:], s[i:])
					s[i].High = lo
					s[j].Low = hi
					*set = s
				}
			} else {
				s[i].High = lo
			}
		} else {
			if hi < s[i].High {
				s[i].Low = hi
			} else {
				s = append(s[:i], s[j:]...)
				*set = s
			}
		}

		return
	}

	// Case 5.

	if lo > s[i].Low {
		s[i].High = lo
		i++
	}

	if hi < s[j-1].High {
		s[j-1].Low = hi
		j--
	}

	s = append(s[:i], s[j:]...)
	*set = s
}

// Contains reports whether or not set contains a single element.
func (set RangeSet[E]) Contains(e E) bool {
	return set.ContainsRange(e, e+1)
}

// ContainsRange reports whether or not set contains every element in range
// [lo, hi).
func (set RangeSet[E]) ContainsRange(lo, hi E) bool {
	i := sort.Search(len(set), func(i int) bool { return set[i].High > lo })
	return i < len(set) && set[i].Low <= lo && hi <= set[i].High && lo < hi
}

// Difference returns the subset of set that having all elements in other
// excluded.
func (set RangeSet[E]) Difference(other RangeSet[E]) RangeSet[E] {
	return set.Intersection(other.Complement())
}

// Equal reports whether or not set is identical to other.
func (set RangeSet[E]) Equal(other RangeSet[E]) bool {
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
func (set RangeSet[E]) Extent() Range[E] {
	if len(set) == 0 {
		return Range[E]{}
	}

	return Range[E]{
		Low:  set[0].Low,
		High: set[len(set)-1].High,
	}
}

// IsSubsetOf reports whether or not other contains every element in set.
func (set RangeSet[E]) IsSubsetOf(other RangeSet[E]) bool {
	for _, r := range set {
		if !other.ContainsRange(r.Low, r.High) {
			return false
		}
	}

	return true
}

// Count returns the number of element in set.
func (set RangeSet[E]) Count() uint64 {
	var count uint64

	for _, r := range set {
		count += uint64(r.High - r.Low)
	}

	return count
}
