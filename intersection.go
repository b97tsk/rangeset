package rangeset

import "sort"

// Intersection returns the intersection of set and other.
func (set RangeSet) Intersection(other RangeSet) RangeSet {
	return IntersectionBuffer(set, other, nil)
}

// Intersection returns the intersection of zero or more sets.
func Intersection(sets ...RangeSet) RangeSet {
	return combine(IntersectionBuffer, sets...)
}

// IntersectionBuffer returns the intersection of s1 and s2, uses buffer as
// its initial backing storage.
func IntersectionBuffer(s1, s2, buffer RangeSet) RangeSet {
	result := buffer[:0]

	for {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}

		if len(s2) == 0 {
			break
		}

		r := s2[0]
		i := sort.Search(len(s1), func(i int) bool { return s1[i].High > r.Low })
		t := s1[i:]
		j := i + sort.Search(len(t), func(i int) bool { return t[i].Low >= r.High })

		if i < j {
			start := len(result)
			result = append(result, s1[i:j]...)

			if r0 := &result[start]; r0.Low < r.Low {
				r0.Low = r.Low
			}

			if r1 := &result[len(result)-1]; r1.High > r.High {
				r1.High = r.High
			}

			j--
		}

		s1, s2 = s1[j:], s2[1:]
	}

	return result
}
