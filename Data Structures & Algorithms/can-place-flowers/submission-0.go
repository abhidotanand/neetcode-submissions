func canPlaceFlowers(flowerbed []int, n int) bool {
    var l int = len(flowerbed)
	var cnt int

	if n == 0 {
		return true
	}

	if l == 1 && flowerbed[0] == 0 {
		return true
	}

	if l > 1 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		cnt++
	}
	
	if l > 1 && flowerbed[l-1] == 0 && flowerbed[l-2] == 0 {
		flowerbed[l-1] = 1
		cnt++
	}

	for i := 1; i < l-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
			flowerbed[i] = 1
			cnt++
		}
	}

	if cnt >= n {
		return true
	}

	return false
}