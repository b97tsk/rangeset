package rangeset

import (
	"sort"

	. "golang.org/x/exp/constraints"
)

// SymmetricDifference returns the symmetric difference of two sets.
func SymmetricDifference[E Integer](s1, s2 RangeSet[E]) RangeSet[E] {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	if len(s1) == 0 {
		return nil
	}

	set := make(RangeSet[E], len(s1), len(s1)+len(s2))
	copy(set, s1)

	for _, r := range s2 {
		symmetricDifferenceRange(&set, r.Low, r.High)
	}

	return set
}

func symmetricDifferenceRange[E Integer](set *RangeSet[E], lo, hi E) {
	s := *set

	i := sort.Search(len(s), func(i int) bool { return s[i].High > lo })
	// j := sort.Search(len(s), func(i int) bool { return s[i].Low >= hi })

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
	j := i + sort.Search(len(t), func(i int) bool { return t[i].Low >= hi })

	if i == j { // Case 1, 2 and 3.
		set.AddRange(lo, hi)
		return
	}

	// Case 4 and 5.

	if i > 0 && lo == s[i-1].High {
		lo = s[i].Low
		s[i-1].High = lo
	}

	if j < len(s) && hi == s[j].Low {
		hi = s[j-1].High
		s[j].Low = hi
	}

	loSame := lo == s[i].Low
	hiSame := hi == s[j-1].High

	if loSame && hiSame {
		for ; i < j-1; i++ {
			s[i] = Range[E]{s[i].High, s[i+1].Low}
		}

		s = append(s[:i], s[j:]...)
		*set = s

		return
	}

	if loSame {
		for ; i < j-1; i++ {
			s[i] = Range[E]{s[i].High, s[i+1].Low}
		}

		if hi < s[i].High {
			s[i].Low = hi
		} else {
			s[i].Low, s[i].High = s[i].High, hi
		}

		return
	}

	if hiSame {
		if lo > s[i].Low {
			s[i].High, lo = lo, s[i].High
			i++
		}

		for ; i < j; i++ {
			s[i].Low, s[i].High, lo = lo, s[i].Low, s[i].High
		}

		return
	}

	if lo < hi {
		s = append(s, Range[E]{})
		copy(s[j:], s[j-1:])
		*set = s

		if lo > s[i].Low {
			s[i].High, lo = lo, s[i].High
			i++
		}

		for ; i < j; i++ {
			s[i].Low, s[i].High, lo = lo, s[i].Low, s[i].High
		}

		if hi < s[i].High {
			s[i].Low = hi
		} else {
			s[i].Low, s[i].High = s[i].High, hi
		}
	}
}
