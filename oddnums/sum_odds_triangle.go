package oddnums

func RowSumOddNumbers(n int) int {
	// count order 'ord'
	// 'ord' stands for combination number (n over 2) -> (n*(n-1)) / 2
	// odd numbers in the nth row are positioned in the triangle: <ord+1, ord+n>
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	ord := (n * (n - 1)) / 2

	// Concrete Digits in the nth row are in the interval: <2*(ord+1)-1, 2*(ord+n)-1>
	sum := 0
	for i := 1; i <= n; i++ {
		sum += 2*(ord+i) - 1
	}
	return sum
}
