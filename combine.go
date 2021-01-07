package rangeset

func combine(
	op func(s1, s2, buffer RangeSet) RangeSet,
	sets ...RangeSet,
) RangeSet {
	if len(sets) == 0 {
		return nil
	}

	var r1, r2 RangeSet

	r1 = append(r1, sets[0]...)

	for _, set := range sets[1:] {
		r2 = op(r1, set, r2)
		r1, r2 = r2, r1
	}

	return r1
}
