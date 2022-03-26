package rangeset

import "sort"

// Union returns the union of set and other.
func (set RangeSet[E]) Union(other RangeSet[E]) RangeSet[E] {
	return unionBuffer(set, other, nil)
}

// Union returns the union of zero or more sets.
func Union[E Elem](sets ...RangeSet[E]) RangeSet[E] {
	return combine(unionBuffer[E], sets...)
}

// unionBuffer returns the union of s1 and s2, using buf as its initial
// backing storage.
func unionBuffer[E Elem](s1, s2, buf RangeSet[E]) RangeSet[E] {
	res := buf[:0]

	for {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}

		if len(s2) == 0 {
			return append(res, s1...)
		}

		r := s2[0]
		s2 = s2[1:]

		i := sort.Search(len(s1), func(i int) bool { return s1[i].Low > r.Low })

		if i > 0 && r.Low <= s1[i-1].High {
			r.Low = s1[i-1].Low
			i--
		}

		res = append(res, s1[:i]...)
		s1 = s1[i:]

	Again:
		j := sort.Search(len(s1), func(i int) bool { return s1[i].High > r.High })
		s1 = s1[j:]

		if len(s1) > 0 && r.High >= s1[0].Low {
			r.High = s1[0].High
			s1, s2 = s2, s1[1:]

			goto Again
		}

		res = append(res, r)
	}
}
