package arrays

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return // no merge is needed here
	}
	for i, j := m-1, n-1; j >= 0; {
		// n1 elem is greater, copy it to the end, +1 since both are 0 based indices
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
}
