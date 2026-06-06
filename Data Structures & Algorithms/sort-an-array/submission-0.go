func sortArray(nums []int) []int {
    var n = len(nums)
	var low int = 0
	var high int = n - 1

	mergeSort(nums, low, high)

	return nums
}

func mergeSort(nums []int, low, high int) {
	if low >= high {
		return
	}

	var mid int = (high+low)/2

	mergeSort(nums, low, mid)
	mergeSort(nums, mid+1, high)
	merge(nums, low, mid, high)
}

func merge(nums []int, low, mid, high int) {
	var tmp []int = make([]int, 0)
	var ptr1, ptr2 int = low, mid + 1

	for ptr1 <= mid && ptr2 <= high{
		min_tmp := min(nums[ptr1], nums[ptr2])
		tmp = append(tmp, min_tmp)
		if min_tmp == nums[ptr1]{
			ptr1++
		} else {
			ptr2++
		}
	}

	for ptr2 <= high{
		tmp = append(tmp, nums[ptr2])
		ptr2++
	}

	for ptr1 <= mid{
		tmp = append(tmp, nums[ptr1])
		ptr1++
	}
	j := 0
	for i := low; i <= high; i++ {
		nums[i] = tmp[j]
		j++
	}
}