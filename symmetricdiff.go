package rangeset

import "sort"

// SymmetricDifference returns the symmetric difference of two sets.
func SymmetricDifference(s1, s2 RangeSet) RangeSet {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	if len(s1) == 0 {
		return nil
	}

	set := make(RangeSet, len(s1), len(s1)+len(s2))
	copy(set, s1)

	for _, r := range s2 {
		symmetricDifferenceRange(&set, r.Low, r.High)
	}

	return set
}

func symmetricDifferenceRange(set *RangeSet, low, high int64) {
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
		set.AddRange(low, high)
		return
	}

	// Case 4 and 5.

	lowIdentical := low == s[i].Low
	highIdentical := high == s[j-1].High

	if lowIdentical && highIdentical {
		for ; i < j-1; i++ {
			s[i] = Range{s[i].High, s[i+1].Low}
		}

		s = append(s[:i], s[j:]...)
		*set = s

		return
	}

	if lowIdentical {
		for ; i < j-1; i++ {
			s[i] = Range{s[i].High, s[i+1].Low}
		}

		if high < s[i].High {
			s[i].Low = high
		} else {
			s[i].Low, s[i].High = s[i].High, high
		}

		return
	}

	if highIdentical {
		if low > s[i].Low {
			s[i].High, low = low, s[i].High
			i++
		}

		for ; i < j; i++ {
			s[i].Low, s[i].High, low = low, s[i].Low, s[i].High
		}

		return
	}

	if low < high {
		s = append(s, Range{})
		copy(s[j:], s[j-1:])
		*set = s

		if low > s[i].Low {
			s[i].High, low = low, s[i].High
			i++
		}

		for ; i < j; i++ {
			s[i].Low, s[i].High, low = low, s[i].Low, s[i].High
		}

		if high < s[i].High {
			s[i].Low = high
		} else {
			s[i].Low, s[i].High = s[i].High, high
		}
	}
}
