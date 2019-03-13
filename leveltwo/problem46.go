package leveltwo

func getNextComposite(n int) int {
	n++
	for n > 1 {
		if !IsPrimeInt(int64(n)) && n%2 != 0 {
			return n
		}
		n++
	}
	return n
}

func FindGoldbachConjectureFail() int {
	i := 9
	isGoldbach := true

	for isGoldbach {
		p := 2
		n := 1
		g := p + 2*n*n
		for i > p {
			if g == i {
				isGoldbach = true
				n = 1
				p = 2
				i = getNextComposite(i)
			} else if g < i {
				n++
				g = p + 2*n*n
			} else if g > i {
				n = 1
				p = findNextPrime(p)
				g = p + 2*n*n
			}
		}
		if i < p {
			isGoldbach = false
		}
	}
	return i
}
