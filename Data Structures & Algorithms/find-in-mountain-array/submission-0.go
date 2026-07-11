/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    var n int = mountainArr.length()
	var left, right = 0, n-1
	var tip int

	for left <= right {
		mid := (left+right)/2

		midEle := mountainArr.get(mid)
		prevEle := mountainArr.get(mid-1)
		nextEle := mountainArr.get(mid+1)
		
		if midEle > prevEle && midEle > nextEle {
			tip = mid
			break
		} else if midEle < prevEle && midEle > nextEle {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left = 0
	right = tip

	for left <= right {
		mid := (left+right)/2

		midEle := mountainArr.get(mid)

		if midEle == target {
			return mid
		} else if midEle > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left = tip + 1
	right = n-1

	for left <= right {
		mid := (left+right)/2

		midEle := mountainArr.get(mid)

		if midEle == target {
			return mid
		} else if midEle > target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
