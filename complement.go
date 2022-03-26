package rangeset

// Complement returns the inverse of set.
//
// Complement of an empty set is the return value of Universal[E](), which
// contains every E except one, the maximum value of E.
func (set RangeSet[E]) Complement() RangeSet[E] {
	if len(set) == 0 {
		return Universal[E]()
	}

	return complement(set)
}

func complement[E Elem](set RangeSet[E]) RangeSet[E] {
	var res RangeSet[E]

	if len(set) > 1 {
		res = make(RangeSet[E], 0, len(set)+1) // Pre-allocation.
	}

	r0 := set[0]

	if r0.Low > minOf[E]() {
		res = append(res, Range[E]{minOf[E](), r0.Low})
	}

	lo := r0.High

	for _, r := range set[1:] {
		res = append(res, Range[E]{lo, r.Low})
		lo = r.High
	}

	if lo < maxOf[E]() {
		res = append(res, Range[E]{lo, maxOf[E]()})
	}

	return res
}
