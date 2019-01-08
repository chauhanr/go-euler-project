package levelone

func SumEvenFibSeries(limit int64) int64 {
	var p1, p2, n, sum int64
	p1 = 1
	p2 = 2
	sum = 0
	n = 0

	for p2 <= limit {
		if p2%2 == 0 {
			sum += p2
		}
		n = nextElement(p1, p2)
		p1 = p2
		p2 = n
	}

	return sum
}

func nextElement(a, b int64) int64 {
	return a + b
}
