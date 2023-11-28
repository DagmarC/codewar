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

func WhoEatsWhoLL(zoo string) []string {

	var result []string
	result = append(result, zoo)

	// eaten animals are in space separated values
	animals := strings.Split(zoo, ",")
	list := NewLinkedlist[string]()
	for _, animal := range animals {
		list.InsertLast(strings.TrimSpace(strings.ToLower(animal)))
	}
	eatingLL(&result, list)

	result = append(result, fmt.Sprint(list))
	fmt.Println(result)
	return result
}

func eatingLL(result *[]string, list *Linkedlist[string]) {
	animal := list.head
	for animal != nil {
		// LEFT
		if animal.prev != nil {

			if canEat(animal.key, animal.prev.key) {
				tmpAnimal := animal.prev.prev // if animal eats to the left, continue with the leftmost animal again, if leftmost is NIL go right
				// eat animal
				*result = append(*result, fmt.Sprintf("%s eats %s", animal.key, animal.prev.key))

				list.DeleteNode(animal.prev)
				if tmpAnimal != nil {
					animal = tmpAnimal // continue to the right (left was already checked before)
				}
			}
		}

		// RIGHT
		// animal eats to its right until nothing else can be eaten by him
		for animal.next != nil && canEat(animal.key, animal.next.key) {
			*result = append(*result, fmt.Sprintf("%s eats %s", animal.key, animal.next.key))
			list.DeleteNode(animal.next)
		}
		animal = animal.next
	}
}

