// условие задачи - https://leetcode.com/problems/merge-sorted-array/description/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	ptr1 := m - 1
	ptr2 := n - 1
	i := m + n - 1

	for ; ptr2 >=0 && ptr1 >= 0; i-- {
		if nums1[ptr1] < nums2[ptr2] {
			nums1[i] = nums2[ptr2]
			ptr2--
		} else {
			nums1[i] = nums1[ptr1]
			ptr1--
		}
	}
	for ptr2 >= 0 {
		nums1[i] = nums2[ptr2]
		ptr2--
		i--
	}
}