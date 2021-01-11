package main

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	len1 := m - 1
	len2 := n - 1
	del := m + n - 1

	for len2 >= 0 {
		if len1 < 0 {
			for len2 >= 0 {
				nums1[del] = nums2[len2]
				del--
				len2--
			}
		} else if nums1[len1] >= nums2[len2] {
			nums1[del] = nums1[len1]
			del--
			len1--
		} else {
			nums1[del] = nums2[len2]
			del--
			len2--
		}
	}
	return nums1
}
