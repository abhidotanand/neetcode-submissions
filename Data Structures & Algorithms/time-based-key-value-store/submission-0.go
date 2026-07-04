type TimeMap struct {
	m map[string][]struct{
		timestamp int
		value string
	}
}

func Constructor() TimeMap {
	return TimeMap{
		make(map[string][]struct{
			timestamp int
			value string
		}),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], struct{
		timestamp int
		value string
	}{timestamp, value})

}

func (this *TimeMap) Get(key string, timestamp int) string {
	if values, ok := this.m[key]; ok {
		var n int = len(values)
		if n == 1 {
			if values[0].timestamp > timestamp {
				return ""
			} else {
				return values[0].value
			}
		}

		if timestamp < values[0].timestamp {
			return ""
		} else if timestamp > values[n-1].timestamp {
			return values[n-1].value
		}

		var left, mid, right int = 0, -1, n-1

		for left <= right {
			mid = (left+right)/2

			if values[mid].timestamp == timestamp {
				return values[mid].value
			}

			if values[mid].timestamp < timestamp {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return values[mid-1].value
	}
	return ""
}
