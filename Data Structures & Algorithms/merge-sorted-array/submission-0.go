func merge(nums1 []int, m int, nums2 []int, n int) {
	var tmp int

	for i := m; i < m+n; i++ {
		nums1[i] = nums2[tmp]
		tmp++
	}

	for i := m; i < m+n; i++ {
		for j := i; j > 0; j-- {
			if nums1[j-1] > nums1[j]{
				nums1[j-1], nums1[j] = nums1[j], nums1[j-1]
			} else {
				break
			}
		}
	}
}
