package main

import (
	"fmt"
)

const (
	digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ._-~"
)

func FormatNumber(v int64, base int) (string, error) {
	if base < 2 || base > len(digits) {
		return "", fmt.Errorf("illegal base %d out of range 2 to %d\n", base, len(digits))
	}

	neg := v < 0
	u := uint64(v)
	if neg {
		u = uint64(-u)
	}

	var a [64 + 1]byte // space for 64bit value in base 2, plus sign bit
	i := len(a) - 1

	b := uint64(base)
	for ; u >= b; i-- {
		q := u / b
		a[i] = digits[uintptr(u - (q * b))]
		u = q
	}

	// now, u < base
	a[i] = digits[uintptr(u)]

	if neg {
		i--
		a[i] = '-'
	}

	return string(a[i:]), nil
}
