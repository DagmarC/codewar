package dynamicproblems

func minFlipsMonoIncr(s string) int {
	var oneCount int
	var res int

	for _, b := range s {
		if b == '1' {
			oneCount++
		} else {
			res = min(oneCount, res+1)
		}
	}
	return res
}
