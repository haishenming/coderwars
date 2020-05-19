package kata

func RowSumOddNumbers(n int) int {
	total := 0
	n = n - 1
	r0 := 1
	for i := 0; i < n; i++ {
		r0 = r0 + 2*(i+1)
	}
	total = r0
	for i := 0; i < n; i++ {
		r0 += 2
		total += r0
	}

	return total
}
