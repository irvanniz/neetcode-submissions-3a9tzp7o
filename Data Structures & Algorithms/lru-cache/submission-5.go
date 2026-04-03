import "slices"

type LRUCache struct {
    data map[int]int
	capacity int
	lruKeys []int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		data: make(map[int]int, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
    v, ok := this.data[key]
	if !ok {
		return -1
	}

	// adjust lru
	lruKeys := slices.DeleteFunc(this.lruKeys, func(n int) bool {
		return n == key
	})
	this.lruKeys = append(lruKeys, key)

	fmt.Printf("[get] key:%d | lru:%+v \n\n", key, this.lruKeys)
	return v
}

func (this *LRUCache) Put(key int, value int) {
    _, isExist := this.data[key]
	this.data[key]=value

	if isExist {
		this.lruKeys = slices.DeleteFunc(this.lruKeys, func(n int) bool {
			return n == key
		})
	}
	this.lruKeys = append(this.lruKeys, key)

	// remove lru key if exceeds
	if len(this.data) > this.capacity {
		delete(this.data, this.lruKeys[0])
		this.lruKeys = slices.DeleteFunc(this.lruKeys, func(n int) bool {
			return n == this.lruKeys[0]
		})
	}

	fmt.Printf("[put] key:%d | lru:%+v, map:%+v \n\n", key, this.lruKeys, this.data)
}
