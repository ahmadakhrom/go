package chicken

func Years(n int) int {
	return n * 7
}

func YearsTwo(n int) int {
	a := 0
	for i := 0; i < n; i++ {
		a += 7
	}
	return a
}
