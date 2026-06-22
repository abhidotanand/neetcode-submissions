import "slices"
func carFleet(target int, position []int, speed []int) int {
	var n int = len(position)
	var cars [][]int = make([][]int, n)
	var tmp []float64 = make([]float64, 0)
	var ans int = 1
	var ttr float64

	for i := 0; i < n; i++ {
		cars[i] = []int{position[i], speed[i]}
	}

	slices.SortFunc(cars, func(a, b []int) int {
		if a[0] > b[0] {
			return -1
		} else if a[0] == b[0] {
			return 0
		}
		return 1
	})

	for i := 0; i < n; i++ {
		ttr = (float64(target) - float64(cars[i][0]))/float64(cars[i][1])
		if len(tmp) == 0 {
			tmp = append(tmp, ttr)
		} else {
			if tmp[len(tmp)-1] >= ttr {
				tmp = append(tmp, tmp[len(tmp)-1])
			} else {
				tmp = append(tmp, ttr)
			}
		}
	}

	for i := 0; i < n-1; i++ {
		if tmp[i] != tmp[i+1] {
			ans++
		}
	}

	return ans
}
