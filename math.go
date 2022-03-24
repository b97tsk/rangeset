package rangeset

import (
	"unsafe"

	. "golang.org/x/exp/constraints"
)

func maxOf[E Integer]() E {
	return ^minOf[E]()
}

func minOf[E Integer]() E {
	if ^E(0) > 0 {
		return 0
	}

	return E(1) << (unsafe.Sizeof(E(0))*8 - 1)
}
