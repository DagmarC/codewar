package slidingwindow

import (
	"errors"
)

type deque []int

func NewDeque() *deque {
	d := make(deque, 0)
	return &d
}

func (d *deque) Len() int {
	return len(*d)
}

func (d *deque) Empty() bool {
	return len(*d) == 0
}

func (d *deque) Get(i int) int {
	return (*d)[i]
}

func (d *deque) First() int {
	return (*d)[0]
}

func (d *deque) Last() int {
	return (*d)[len(*d)-1]
}

func (d *deque) Pop() (int, error) {
	if len(*d) == 0 {
		return -1, errors.New("empty queue")
	}
	el := (*d)[0]
	*d = (*d)[1:]
	return el, nil
}

func (d *deque) PopLast() (int, error) {
	if len(*d) == 0 {
		return -1, errors.New("empty queue")
	}
	el := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return el, nil
}

func (d *deque) Push(n int) {
	*d = append(*d, n)
}

func maxSlidingWindow(nums []int, k int) []int {
	recentlySeenMax := NewDeque()
	maximums := make([]int, 0, len(nums))

	for i := range nums {
		removeOldIndex(recentlySeenMax, i, k)
		updateQueue(recentlySeenMax, nums, i)
		addTomaximums(recentlySeenMax, &maximums, nums, i, k)
	}
	return maximums
}

func removeOldIndex(recentlySeenMax *deque, i, k int) {
	if recentlySeenMax.Len() != 0 && i-k == recentlySeenMax.First() {
		recentlySeenMax.Pop()
	}
}

func updateQueue(recentlySeenMax *deque, nums []int, i int) {
	for recentlySeenMax.Len() != 0 && nums[(*recentlySeenMax).Last()] < nums[i] {
		recentlySeenMax.PopLast()
	}
	recentlySeenMax.Push(i)
}

func addTomaximums(recentlySeenMax *deque, maximums *[]int, nums []int, i, k int) {
	if i >= k-1 {
		*maximums = append(*maximums, nums[recentlySeenMax.First()])
	}
}
