package zoo

import (
	"fmt"
	"strings"
)

/*
antelope eats grass
big-fish eats little-fish
bug eats leaves
bear eats big-fish
bear eats bug
bear eats chicken
bear eats cow
bear eats leaves
bear eats sheep
chicken eats bug
cow eats grass
fox eats chicken
fox eats sheep
giraffe eats leaves
lion eats antelope
lion eats cow
panda eats leaves
sheep eats grass
*/

// NOTES: doubly connected LL, look to the LEFT (prev) first and then to the RIGHT
// current animal (cur) eats first left then right
// before eating the animal to the left, set prev animal of left animal cur.prev.prev
// as the one that will be eating next (the leftmost animal eats first)
// in casse of eating right animal, leave current animal until it finds no more animals to eat next to him
// when current animal cant eat anything, set current animal to animal.next
// use map: keys animals that eats and values []string of animals that are eaten by the key animal
// split input string and save it to the LL

var eaters = map[string]string{
	"antelope": "grass",
	"big-fish": "little-fish",
	"bug":      "leaves",
	"bear":     "big-fish bug chicken cow leaves sheep",
	"chicken":  "bug",
	"cow":      "grass",
	"fox":      "chicken sheep",
	"giraffe":  "leaves",
	"lion":     "antelope cow",
	"panda":    "leaves",
	"sheep":    "grass",
}

func WhoEatsWho(zoo string) []string {

	var result []string
	result = append(result, zoo)

	// eaten animals are in space separated values
	animals := strings.Split(zoo, ",")

	eating(&result, &animals)

	result = append(result, strings.Join(animals, ","))
	return result
}

func eating(result *[]string, list *[]string) {
	i := 0
	for len(*list) != 0 && i < len(*list) {
		// LEFT
		if i > 0 && i < len(*list) {
			if canEat((*list)[i], (*list)[i-1]) {
				*result = append(*result, fmt.Sprintf("%s eats %s", (*list)[i], (*list)[i-1])) // eat animal to its left
				*list = append((*list)[:i-1], (*list)[i:]...) // delete left animal
				i--                                           // eater animal's index decreased, cuz previous animal was eaten
				if i >= 1 {
					i = i - 1 // Leftmost eats first
				}
			}
		}

		// RIGHT
		// animal eats to its right until nothing else can be eaten by him
		for len(*list) != 0 && i < len(*list)-1 && canEat((*list)[i], (*list)[i+1]) {
			*result = append(*result, fmt.Sprintf("%s eats %s", (*list)[i], (*list)[i+1]))
			if i < len(*list)-2 {
				*list = append((*list)[:i+1], (*list)[i+2:]...) // delete right animal
			} else {
				// right animal is the last element
				*list = (*list)[:i+1] // delete the right animal
			}
		}
		i++
	}
}

func canEat(eater, eaten string) bool {
	return strings.Contains(eaters[eater], eaten)
}
