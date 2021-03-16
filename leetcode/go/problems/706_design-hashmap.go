package problems

// 拉链法:
// 定义了一个比较小的数组，然后使用 hash 方法来把求出 key 应该出现在数组中的位置；
// 但是由于不同的 key 在求完 hash 之后，可能会存在碰撞冲突，
// 所以数组并不直接保存元素，而是每个位置都指向了一条链表（或数组）用于存储元素。

// 查找一个 key 的时候需要两个步骤：
// 求hash到数组中的位置；
// 在链表中遍历找key。

// 不定长的拉链数组是说拉链会根据分桶中的 key 动态增长
// 把分桶数去了 1009，是因为 1009 是大于 1000 的第一个质数
// https://leetcode-cn.com/problems/design-hashmap/solution/xiang-jie-hashmap-de-she-ji-zai-shi-jian-85k9/

type MyHashMap struct {
	buckets int
	table [][]listnode
}

type listnode struct {
	key   int
	value int
}

/** Initialize your data structure here. */
func Constructor706() MyHashMap {
	return MyHashMap{
		buckets:1009,
		table:make([][]listnode,1009),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	k:=key%this.buckets
	list:=&this.table[k]
	for i, v := range *list {
		if(v.key== key){
			(*list)[i].value=value
			return
		}
	}
	// key不存在直接添加
	*list = append(*list,listnode{key,value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	k:=key%this.buckets
	list:=this.table[k]
	for _, v := range list {
		if(v.key== key){
			return v.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	k:=key%this.buckets
	list:=&this.table[k]
	for i, v := range *list {
		if(v.key== key){
			*list = append((*list)[:i],(*list)[i+1:]...)
			return
		}
	}
	// key不存在，什么都不需要做
	
}

