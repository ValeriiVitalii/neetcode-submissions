type MyHashMap struct {
	Capacity int
	Pairs [][]*Pair
}

func Constructor() MyHashMap {
	const capacity = 1000

    return MyHashMap{
		Capacity: capacity,
		Pairs: make([][]*Pair, capacity),
	}
}

func (this *MyHashMap) Put(key int, value int) {
    index := this.Hash(key)

	for _, pair := range this.Pairs[index] {
		if pair.Key == key {
			pair.Value = value

			return
		}
	}

	this.Pairs[index] = append(this.Pairs[index], newPair(key, value))
}

func (this *MyHashMap) Get(key int) int {
    index := this.Hash(key)

	for _, pair := range this.Pairs[index] {
		if pair.Key == key {
			return pair.Value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
    index := this.Hash(key)

	for i, pair := range this.Pairs[index] {
		if pair.Key == key {
			this.Pairs[index][i] = this.Pairs[index][len(this.Pairs[index]) - 1]
			this.Pairs[index] = this.Pairs[index][:len(this.Pairs[index]) - 1]

			return
		}
	}
}

func (this *MyHashMap) Hash(key int) int {
	return key % this.Capacity
}

type Pair struct {
	Key, Value int
}

func newPair(key, value int) *Pair {
	return &Pair{Key: key, Value: value}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */