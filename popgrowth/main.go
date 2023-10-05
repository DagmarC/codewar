package popgrowth

import "math"

// NbYear function should return n number of entire years needed to get a population greater or equal to p.
func NbYear(p0 int, percent float64, aug int, p int) int {
	var n int // result, number of entire years needed to get a population greater than p
	popgrowth := float64(p0)
	for math.Floor(popgrowth) <= float64(p) {
		popgrowth = float64(popgrowth) + float64(popgrowth)*(percent/100.0) + float64(aug)
		n++
	}
	return n
}
