package rangeset

import "sort"

// Intersect returns the intersection of set and other.
func (set RangeSet) Intersect(other RangeSet) RangeSet {
	return IntersectBuffer(set, other, nil)
}

// Intersect returns the intersection of zero or more sets.
func Intersect(sets ...RangeSet) RangeSet {
	if len(sets) == 0 {
		return nil
	}

	var r1, r2 RangeSet

	r1 = make(RangeSet, len(sets[0]))
	copy(r1, sets[0])

	for _, set := range sets[1:] {
		r2 = IntersectBuffer(r1, set, r2)
		r1, r2 = r2, r1
	}

	return r1
}

// IntersectBuffer returns the intersection of s1 and s2, uses buffer as its
// initial backing storage.
func IntersectBuffer(s1, s2, buffer RangeSet) RangeSet {
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
