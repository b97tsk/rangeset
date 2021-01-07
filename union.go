package rangeset

import "sort"

// Union returns the union of set and other.
func (set RangeSet) Union(other RangeSet) RangeSet {
	return UnionBuffer(set, other, nil)
}

// Union returns the union of zero or more sets.
func Union(sets ...RangeSet) RangeSet {
	return combine(UnionBuffer, sets...)
}

// UnionBuffer returns the union of s1 and s2, uses buffer as its initial
// backing storage.
func UnionBuffer(s1, s2, buffer RangeSet) RangeSet {
	result := buffer[:0]

	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	if len(s2) == 0 {
		return append(result, s1...)
	}

	r := s2[0]
	s2 = s2[1:]

	for {
		i := sort.Search(len(s1), func(i int) bool { return s1[i].Low > r.Low })
		j := sort.Search(len(s1), func(i int) bool { return s1[i].High > r.High })

		if i > 0 && r.Low <= s1[i-1].High {
			r.Low = s1[i-1].Low
			i--
		}

		if j < len(s1) && r.High >= s1[j].Low {
			r.High = s1[j].High
			j++
		}

		result = append(result, s1[:i]...)

		if len(s2) == 0 {
			result = append(result, r)
			result = append(result, s1[j:]...)

			break
		}

		s1, s2 = s2, s1[j:]
	}

	return result
}
