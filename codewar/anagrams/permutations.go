package anagrams

import (
	"fmt"
	"sort"
)

var visited map[string]struct{} = make(map[string]struct{})

func CountPermOrdered(word string) int {
	chars := []rune(word)
	sort.Slice(chars, func(i, j int) bool { //sort the string using the function
		return chars[i] < chars[j]
	})
	res := recPosition(chars, word, 0, 0)
	fmt.Println(visited)
	return res

}

func recPosition(letters []rune, word string, fix int, pos int) int {
	if fix == len(word)-1 {
		if _, ok := visited[string(letters)]; !ok {
			visited[string(letters)] = struct{}{}
			pos++
		}
	}
	for i := fix; i < len(word); i++ {
		if i != fix && letters[fix] == letters[i] {
			continue
		}
		swap(fix, i, letters)
		pos = recPosition(letters, word, fix+1, pos)
		swap(i, fix, letters)

	}
	return pos
}

func swap(i, j int, letters []rune) {
	letters[i], letters[j] = letters[j], letters[i]
}
