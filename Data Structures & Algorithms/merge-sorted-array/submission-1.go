func merge(nums1 []int, m int, nums2 []int, n int) {
	var last, rear1, rear2 int = m+n-1, m-1, n-1

	for rear2 >= 0 && rear1 >=0 {
		if nums1[rear1] > nums2[rear2] {
			nums1[last] = nums1[rear1]
			rear1--
		} else {
			nums1[last] = nums2[rear2]
			rear2--
		}
		last--
	}
	if rear1 <= 0 {
		for rear2 >= 0 {
			nums1[rear2] = nums2[rear2]
			rear2--
		}
	}
}
