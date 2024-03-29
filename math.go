package rangeset

import "unsafe"

func maxOf[E Elem]() E {
	return ^minOf[E]()
}

func minOf[E Elem]() E {
	if ^E(0) > 0 {
		return 0
	}

	return E(1) << (unsafe.Sizeof(E(0))*8 - 1)
}
