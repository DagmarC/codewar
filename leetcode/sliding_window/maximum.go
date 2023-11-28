package slidingwindow

import (
	"slices"
)

func maxSlidingWindow(nums []int, k int) []int {
	maximums := make([]int, 0)

	if len(nums) == 0 {
		return []int{}
	}

	var max = slices.Max(nums[:k])
	maximums = append(maximums, max)

	for i := 1; i <= len(nums)-k; i++ {
		if max == nums[i-1] {
			max = slices.Max(nums[i:k+i])
		}
		if max < nums[k+i-1] { // next element added to the view
			max = nums[k+i-1]
		}
		maximums = append(maximums, max)
	}
	return maximums
}
