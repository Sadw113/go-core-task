package main

import "fmt"

const ArraySize = 7

type StringIntMap struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func (h *StringIntMap) Add(key string, value int) {
	index := hash(key)
	h.array[index].insert(key, value)
}

func (h *StringIntMap) Exists(key string) (int, bool) {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *StringIntMap) Remove(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (h *StringIntMap) Copy() map[string]int {
	res := make(map[string]int)
	for _, val := range h.array {
		if val.head != nil {
			head := val.head
			for head != nil {
				key := head.key
				res[key] = head.value
				head = head.next
			}
		}
	}
	return res

}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func (b *bucket) insert(key string, value int) {
	_, ok := b.search(key)
	if !ok {
		newNode := &bucketNode{key: key, value: value}
		newNode.next = b.head
		b.head = newNode
	} else {
		currNode := b.head
		for currNode.key != key {
			currNode = currNode.next
		}
		currNode.value = value
	}
}

func (b *bucket) search(key string) (int, bool) {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return currNode.value, true
		}
		currNode = currNode.next
	}
	return 0, false
}

func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next.key != key {
		prevNode = prevNode.next
	}
	prevNode.next = prevNode.next.next
}

func main() {
	hashTable := &StringIntMap{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}

	fmt.Println("Hash Table Created")
	hashTable.Add("name", 123)
	hashTable.Add("age", 456)
	val, ok := hashTable.Exists("name")
	if !ok {
		fmt.Println("Ошибка")
	} else {
		fmt.Println("name:", val)
	}
	val, ok = hashTable.Exists("age")
	if !ok {
		fmt.Println("Ошибка")
	} else {
		fmt.Println("name:", val)
	}

	hashTable.Remove("name")

	val, ok = hashTable.Exists("name")
	if !ok {
		fmt.Println("Нет такого значения")
	} else {
		fmt.Println("name:", val)
	}

	newMap := hashTable.Copy()
	fmt.Println(newMap)

	hashTable.Add("name", 4)
	newMap = hashTable.Copy()
	fmt.Println(newMap)
}
