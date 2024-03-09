package dynamicprog

import (
	"bytes"
)

// LCS finds the longest common subsequence in both strings and returns it. Dynamic Programming - memoization is applied here.
func LCS(s1, s2 string) string {
	m := len(s1)
	n := len(s2)
	table := init2DTable(m+1, n+1)

	LCSTable(s1, s2, m, n, &table)

	return getLCS(s1, s2, m, n, &table)
}

func LCSTable(s1, s2 string, m, n int, table *[][]int){
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				(*table)[i][j] = 0
				continue
			}
			if s1[i-1] == s2[j-1] {
				(*table)[i][j] = (*table)[i-1][j-1] + 1
			} else {
				(*table)[i][j] = max((*table)[i][j-1], (*table)[i-1][j])
			}
		}
	}
}

func getLCS(s1, s2 string, i, j int, table *[][]int) string {
	var lcs bytes.Buffer
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs.WriteByte(s1[i-1])
			i--
			j--
			continue
		}
		if (*table)[i-1][j] >= (*table)[i][j-1] {
			i--
		} else {
			j--
		}
	}
	return reversed(lcs.String())
}

func init2DTable(m, n int) [][]int {
	table := make([][]int, m)
	for i := range table {
		table[i] = make([]int, n)
	}
	return table
}

func reversed(s string) string {
	var reversed bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		reversed.WriteByte(s[i])
	}
	return reversed.String()
}
