package main

// 感谢 https://leetcode-cn.com/problems/merge-sorted-array/solution/88-by-ikaruga/
// 提供从后往前排序的思路，从大往小排
func merge(nums1 []int, m int, nums2 []int, n int) {
	o := m + n - 1
	m--
	n--
	for o >= 0 {
		if m >= 0 && n >= 0 {
			if nums1[m] > nums2[n] {
				nums1[o] = nums1[m]
				m--
			} else {
				nums1[o] = nums2[n]
				n--
			}
		} else {
			for n >= 0 {
				nums1[o] = nums2[n]
				n--
				o--
			}
		}
		o--
	}
}
