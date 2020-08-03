package lru2

type LRUCache struct {
	Capacity int
	Mp map[int]int
	List []int
}

//Constructor 构造函数
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Mp: make(map[int]int, 0),
		List: make([]int, 0),
	}
}

//DeleteKey 删除map中的对应key
func DeleteKey(list []int, key int) []int {
	for i := 0; i < len(list); i++ {
		if list[i] == key {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list[1:]
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Mp[key]; ok {
		this.List = DeleteKey(this.List, key)
		this.List = append(this.List, key)
		return v
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	if _, ok := this.Mp[key]; ok {
		this.Mp[key] = value
		this.List = DeleteKey(this.List, key)
		this.List = append(this.List, key)
		return
	}
	if len(this.Mp) >= this.Capacity {
		delKey := this.List[0]
		this.List = DeleteKey(this.List, key)
		delete(this.Mp, delKey)
	}
	this.Mp[key] = value
	this.List = append(this.List, key)
}