package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := float64(1)

	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		n >>= 1
		x *= x
	}

	return res
}

// 11
// 1111
