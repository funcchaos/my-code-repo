package problems

type MyHashSet struct {
	bitset []uint64
}

/** Initialize your data structure here. */
func Constructor705() MyHashSet {
	// 范围是10^6,一个uint64可以表示64个数，这里直接创建足够大的切片，不做扩容处理了
	return MyHashSet{bitset: make([]uint64,15626)}
}

func (this *MyHashSet) Add(key int) {
	// 计算桶的位置
	l := key / 64
	// 计算桶内位置
	bit := key % 64
	// 二进制指定位置来表示添加的数
	this.bitset[l] = (this.bitset[l]) | (1 << bit)
}

func (this *MyHashSet) Remove(key int) {
	l := key / 64
	bit := key % 64
	this.bitset[l] = (this.bitset[l]) & (^(1 << bit))
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	l := key / 64
	bit := key % 64
	// 判断二进制对应位上是不是1
	return 1 == ((this.bitset[l] >> bit) & 1)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
