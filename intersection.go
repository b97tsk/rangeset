package rangeset

import "sort"

// Intersection returns the intersection of set and other.
func (set RangeSet) Intersection(other RangeSet) RangeSet {
	return intersectionBuffer(set, other, nil)
}

// Intersection returns the intersection of zero or more sets.
func Intersection(sets ...RangeSet) RangeSet {
	return combine(intersectionBuffer, sets...)
}

// intersectionBuffer returns the intersection of s1 and s2, using buffer
// as its initial backing storage.
func intersectionBuffer(s1, s2, buffer RangeSet) RangeSet {
	result := buffer[:0]

	for {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}

		if len(s2) == 0 {
			return result
		}

		r := s2[0]
		s2 = s2[1:]

		i := sort.Search(len(s1), func(i int) bool { return s1[i].High > r.Low })
		s1 = s1[i:]
		j := sort.Search(len(s1), func(i int) bool { return s1[i].Low >= r.High })

		if j > 0 {
			start := len(result)
			result = append(result, s1[:j]...)

			if r0 := &result[start]; r0.Low < r.Low {
				r0.Low = r.Low
			}

			if r1 := &result[len(result)-1]; r1.High > r.High {
				r1.High = r.High
			}

			j--
		}

		s1 = s1[j:]
	}
}
