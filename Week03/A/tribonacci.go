func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}

	p, q, r, s := 0, 1, 1, 2
	for i := 3; i < n; i++ {
		p = q
		q = r
		r = s
		s = p + q + r
	}
	return s

}