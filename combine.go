package rangeset

func combine[E Elem](
	op func(s1, s2, buf RangeSet[E]) RangeSet[E],
	sets ...RangeSet[E],
) RangeSet[E] {
	if len(sets) == 0 {
		return nil
	}

	var r1, r2 RangeSet[E]

	r1 = sets[0]
	r1ReferenceToSets0 := true

	for _, set := range sets[1:] {
		r2 = op(r1, set, r2)

		if r1ReferenceToSets0 {
			r1 = nil
			r1ReferenceToSets0 = false
		}

		r1, r2 = r2, r1
	}

	if r1ReferenceToSets0 {
		// Always return a distinct set (unless it's nil).
		r1 = append(RangeSet[E](nil), r1...)
	}

	return r1
}
