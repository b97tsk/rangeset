package rangeset

import "math"

// Complement returns the inverse of set.
func (set RangeSet) Complement() RangeSet {
	if len(set) == 0 {
		return Universal()
	}

	return complementNonEmpty(set)
}

func complementNonEmpty(set RangeSet) RangeSet {
	var result RangeSet

	if len(set) > 1 {
		result = make(RangeSet, 0, len(set)+1) // Pre-allocation.
	}

	r0 := set[0]

	if r0.Low > math.MinInt64 {
		result = append(result, Range{math.MinInt64, r0.Low})
	}

	low := r0.High

	for _, r := range set[1:] {
		result = append(result, Range{low, r.Low})
		low = r.High
	}

	if low < math.MaxInt64 {
		result = append(result, Range{low, math.MaxInt64})
	}

	return result
}
