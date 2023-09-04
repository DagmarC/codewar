package dna

import (
	"strings"
)

//   var complements = map[string]string{"A":"T", "C":"G", "G":"C", "T":"A"}

func DNAStrand(dna string) string {
	// map that keeps track of all complements in DNA
	complements := make(map[rune]rune, 4)
	complements['A'] = 'T'
	complements['T'] = 'A'
	complements['C'] = 'G'
	complements['G'] = 'C'

	var sb strings.Builder
	for _, s := range dna {
		if cp, ok := complements[s]; ok {
			sb.WriteRune(cp)
		}
	}
	return sb.String()
}
