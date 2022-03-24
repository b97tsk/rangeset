package rangeset

import "sort"

// Overlaps reports whether or not the intersection of set and other are
// not empty.
func (set RangeSet[E]) Overlaps(other RangeSet[E]) bool {
	s1, s2 := set, other

	for {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}

		if len(s2) == 0 {
			return false
		}

		r := s2[0]
		s2 = s2[1:]

		i := sort.Search(len(s1), func(i int) bool { return s1[i].High > r.Low })
		s1 = s1[i:]
		j := sort.Search(len(s1), func(i int) bool { return s1[i].Low >= r.High })

		if j > 0 {
			return true
		}

		s1 = s1[j:]
	}
}
