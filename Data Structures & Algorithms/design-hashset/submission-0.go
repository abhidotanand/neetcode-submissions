import (
    "slices"
)

type MyHashSet struct {
	Data []int
}

func Constructor() MyHashSet {
	return MyHashSet{Data: []int{}}
}

func (this *MyHashSet) Add(key int) {
	if !slices.Contains(this.Data, key) {
		this.Data = append(this.Data, key)
	}
	return
}

func (this *MyHashSet) Remove(key int) {
	var n int = len(this.Data)
	var ind int = -1
	for i := 0; i < n; i++ {
		if this.Data[i] == key {
			ind = i
			break
		}
	}
	slices.Delete(this.Data, ind, ind+1)
	return
}

func (this *MyHashSet) Contains(key int) bool {
	if slices.Contains(this.Data, key) {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */