package anagrams

import (
	"math/big"
	"sort"
)

func ListPosition(word string) *big.Int {
	var result = big.NewInt(0)
	m := parse(word) // map of letter and its number of occurrences in given word

	// chars is a sorted sequence of letters obtained from the word
	chars := []rune(word)
	sort.Slice(chars, func(i, j int) bool { //sort the string using the function
		return chars[i] < chars[j]
	})

	calculateRec(&chars, result, word, 0, m)
	return result
}

// calculateRec performs operations needed to have the result of the rank of the word from its permutation
// sol.txt describes the precise algorithm
// Idea is that you go over a sorted chars from the beginning pos=0 (which is a fixes a position) until you find a match
// when match is not found at chars[i] ?= word[pos] you will calculate the permutation with repetition from fixed position len(word)-1-pos

// eg w: BAC, chars: ABC, pos=0, i=0 => chars[i] = A, word[pos] = B (we are looking for letter B in sorted chars)
// A != B no match - calculate permutations A _ _ (letters to word[pos] are fixed so only word[0] is fixed 
// so number of word that can be formed starting with letter A is 2! ... and etc

// when you will find match chars[i] == word[letter], you will delete letter from sorted chars and repeat the process again and go over the sorted chars, 
// calculate number of possible word until match is found
func calculateRec(chars *[]rune, result *big.Int, word string, pos int, m map[rune]int) {
	if len(*chars) == 0 {
		result.Add(result, big.NewInt(1))
		return
	}
	// go over sorted letters in chars 1 by 1 and if letter matched the word[pos]
	// then delete it from map and sorted chars and start over again (rec)
	for i, letter := range *chars {
		if byte(letter) == word[pos] {
			// delete matched letter from slice
			deleteLetter(chars, i, m, letter)
			calculateRec(chars, result, word, pos+1, m) // go from the beginning of chars again at next position
			return
		} else {
			// number of words that can be created from letters (from fixer pos)
			result.Add(result, permRepet(len(word)-1-pos, m))
		}
	}
}

// permRepet calculates permutations with repetition from map m
// map m is map of letter and its count in word to be calculated
func permRepet(n int, m map[rune]int) *big.Int {
	var nf = big.NewInt(int64(n))
	nf.MulRange(1, nf.Int64())

	var divisor = big.NewInt(int64(1))
	for _, v := range m {
		if v > 1 {
			d := big.NewInt(int64(v))
			d.MulRange(1, d.Int64())
			divisor.Mul(divisor, d)
		}
	}
	nf.Div(nf, divisor)
	return nf
}

// deleteLetter deletes letter from ordered string in s and from map subtracts 1 or delete if only one letter has left
func deleteLetter(s *[]rune, i int, m map[rune]int, letter rune) {
	if m[letter] == 1 {
		delete(m, letter)
	} else {
		m[letter]--
	}
	if i == len(*s)-1 {
		*s = (*s)[:i]
	} else {
		*s = append((*s)[:i], (*s)[i+1:]...)
	}
}

// parse creates map of rune/count pairs from string w
func parse(w string) map[rune]int {
	m := make(map[rune]int, 26)
	for _, l := range w {
		m[l]++
	}
	return m
}
