package minpossibleint1505

// The minInteger function (naive solution) in Go takes a string of numbers and an integer k, and rearranges the numbers
// to minimize the value while considering k swaps.
func minInteger(num string, k int) string {
	arr := []byte(num)
	li := 0 // left boundary starts from 0
	for li < len(arr) && k > 0 && k+li > li {
		iMin := findMinInterval(li, k+li+1, &arr)
		for k > 0 && iMin > li {
			arr[iMin], arr[iMin-1] = arr[iMin-1], arr[iMin] // Swtich
			iMin--
			k--
		}
		li++ // move boundaries
	}
	return string(string(arr))
}

// The function `findMinInterval` finds the index of the minimum value in a specified interval of a
// byte array.
func findMinInterval(li int, ri int, arr *[]byte) int {
	iMin := li
	for i := li + 1; i < len(*arr) && i < ri; i++ {
		if (*arr)[i] < (*arr)[iMin] {
			iMin = i
		}
	}
	return iMin
}
