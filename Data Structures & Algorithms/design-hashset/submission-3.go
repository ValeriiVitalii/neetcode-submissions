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

	if !slices.Contains(this.Values[index], key) {
		this.Values[index] = append(this.Values[index], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	index := this.Hash(key)

	for i := 0; i < len(this.Values[index]); i++ {
		if this.Values[index][i] == key {
			this.Values[index][i] = this.Values[index][len(this.Values[index])-1]
			this.Values[index] = this.Values[index][:len(this.Values[index])-1]

			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    index := this.Hash(key)

	return slices.Contains(this.Values[index], key)
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
 