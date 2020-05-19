package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
	if p0 >= p {
		return 0
	}
	
	return NbYear(p0 + int(float64(p0)*percent/100) + aug, percent, aug, p) + 1
}
