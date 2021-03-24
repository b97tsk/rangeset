package rangeset

import "sort"

// Union returns the union of set and other.
func (set RangeSet) Union(other RangeSet) RangeSet {
	return unionBuffer(set, other, nil)
}

// Union returns the union of zero or more sets.
func Union(sets ...RangeSet) RangeSet {
	return combine(unionBuffer, sets...)
}

// unionBuffer returns the union of s1 and s2, using buffer as its initial
// backing storage.
func unionBuffer(s1, s2, buffer RangeSet) RangeSet {
	result := buffer[:0]

	for {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}

		if len(s2) == 0 {
			return append(result, s1...)
		}

		r := s2[0]
		s2 = s2[1:]

		i := sort.Search(len(s1), func(i int) bool { return s1[i].Low > r.Low })

		if i > 0 && r.Low <= s1[i-1].High {
			r.Low = s1[i-1].Low
			i--
		}

		result = append(result, s1[:i]...)
		s1 = s1[i:]

	Again:
		j := sort.Search(len(s1), func(i int) bool { return s1[i].High > r.High })
		s1 = s1[j:]

		if len(s1) > 0 && r.High >= s1[0].Low {
			r.High = s1[0].High
			s1, s2 = s2, s1[1:]

			goto Again
		}

		result = append(result, r)
	}
}
