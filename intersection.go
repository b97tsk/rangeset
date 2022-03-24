package rangeset

import (
	"sort"

	. "golang.org/x/exp/constraints"
)

// Intersection returns the intersection of set and other.
func (set RangeSet[E]) Intersection(other RangeSet[E]) RangeSet[E] {
	return intersectionBuffer[E](set, other, nil)
}

// Intersection returns the intersection of zero or more sets.
func Intersection[E Integer](sets ...RangeSet[E]) RangeSet[E] {
	return combine(intersectionBuffer[E], sets...)
}

// intersectionBuffer returns the intersection of s1 and s2, using buf as
// its initial backing storage.
func intersectionBuffer[E Integer](s1, s2, buf RangeSet[E]) RangeSet[E] {
	res := buf[:0]

	for {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}

		if len(s2) == 0 {
			return res
		}

		r := s2[0]
		s2 = s2[1:]

		i := sort.Search(len(s1), func(i int) bool { return s1[i].High > r.Low })
		s1 = s1[i:]
		j := sort.Search(len(s1), func(i int) bool { return s1[i].Low >= r.High })

		if j > 0 {
			start := len(res)
			res = append(res, s1[:j]...)

			if r0 := &res[start]; r0.Low < r.Low {
				r0.Low = r.Low
			}

			if r1 := &res[len(res)-1]; r1.High > r.High {
				r1.High = r.High
			}

			j--
		}

		s1 = s1[j:]
	}
}
