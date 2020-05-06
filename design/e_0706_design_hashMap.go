package design

/*
Company: Amazon(11), Microsoft(6), Goldman Sachs(5), Oracle(5), Adobe(4)

Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
get(key): Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.

Note:
All keys and values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashMap library.
*/

// MyHashMap ...
type MyHashMap struct {
	Dict []int
}

// Constructor706 ...
func Constructor706() MyHashMap {
	dict := make([]int, 10000, 1000001)
	for i := range dict {
		dict[i] = -1
	}
	return MyHashMap{
		Dict: dict,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	n := len(this.Dict)
	for key >= n {
		n *= 2
	}
	dif := n - len(this.Dict)
	for i := 0; i < dif; i++ {
		this.Dict = append(this.Dict, -1)
	}
	this.Dict[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if key >= len(this.Dict) {
		return -1
	}
	return this.Dict[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if key >= len(this.Dict) {
		return
	}
	this.Dict[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
