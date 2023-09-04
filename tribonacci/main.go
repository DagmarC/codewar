package tribonacci

func Tribonacci(signature [3]float64, n int) []float64 {
	res := make([]float64, 0, n)
	if n == 0 {
		return res
	}
	if n <= 3 {
		return signature[:n]
	} else {
		res = append(res, signature[:3]...)
		recTribonacci(&res, 0, n)
	}
	return res
}

// recTribonacci fills in the res array and stops when the len of the res is equal to the n
// n is at least 4 at the beginning, i should start at 0
func recTribonacci(res *[]float64, i int, n int) {
	if len(*res) == n {
		return
	}
	if n-i <= 0 {
		return // i+2 is out of range
	}
	trib := (*res)[i] + (*res)[i+1] + (*res)[i+2]
	*res = append(*res, trib)
	recTribonacci(res, i+1, n)
}
