import (
	"slices"
)

type MyHashSet struct {
	Capacity int
	Values [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{
		Values: make([][]int, 2),
		Capacity: 2,
	}
}

func (this *MyHashSet) Add(key int) {
	index := this.Hash(key)

	if this.Values[index] != nil {
		if !slices.Contains(this.Values[index], key) {
			this.Values[index] = append(this.Values[index], key)
		}
	} else {
		this.Values[index] = []int{key}
	}
}

func (this *MyHashSet) Remove(key int) {
	index := this.Hash(key)

	for i := 0; i < len(this.Values[index]); i++ {
		if this.Values[index][i] == key {
			this.Values[index][i] = this.Values[index][len(this.Values[index])-1]
			this.Values[index] = this.Values[index][:len(this.Values[index])-1]
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    index := this.Hash(key)

	values := this.Values[index]

	for _, value := range values {
		if key == value {
			return true
		}
	}

	return false
}

func (this *MyHashSet) Hash(key int) int {
	return key % this.Capacity
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 