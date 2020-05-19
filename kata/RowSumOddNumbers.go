package kata

func RowSumOddNumbers(n int) int {
	r0 := 1 + (n-1) * 2
	for i:=0;i<(n-1);i++{
		r0 += 2
	}

	return r0
}
