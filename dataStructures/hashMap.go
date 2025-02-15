package datastructures

import "fmt"

type hNode struct {
	key   string
	value string
	next  *hNode
}

type HashMap struct {
	buckets []*hNode // Array of pointers to Entry structs
	size    int      // Number of buckets
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([]*hNode, size),
		size:    size,
	}
}

func (hTable *HashMap) hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash + int(char)) % hTable.size
	}
	return hash
}

func (hTable *HashMap) PutFunction(key, value string) {
	index := hTable.hashFunction(key);
	node := &hNode{key: key, value: value};

	if hTable.buckets[index] == nil {
		hTable.buckets[index] = node
	} else {
		temp := hTable.buckets[index]
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	}
}

func (hTable *HashMap) GetFunction(key string) string {
	index := hTable.hashFunction(key);
	temp := hTable.buckets[index];

	// if temp.key == key {
	// 	return temp.value
	// } else {
	// 	for temp.key != key {
	// 		temp = temp.next;
	// 	}
	// 	return temp.value
	// }

	for temp != nil {
		if temp.key == key {
			return temp.value
		} else {
			temp = temp.next
		}
	}

	// in case there is no node for the required key
	// empty string as value is returned
	return ""
}

func (hTable *HashMap) DeleteFunction(key string) {
	index := hTable.hashFunction(key);
	temp := hTable.buckets[index];
	var prev *hNode

	for temp != nil {
		if temp.key == key {
			if prev != nil {
				prev.next = temp.next
			} else {
				hTable.buckets[index] = temp.next
			}
			return
		}
		prev = temp
		temp = temp.next
	}
}

func (hTable *HashMap) Display() {
	index := 0;
	for ; index < hTable.size; index++ {
		temp := hTable.buckets[index];
        fmt.Printf("Bucket %d: ", index);
        for temp != nil {
            fmt.Printf("(%s, %s) -> ", temp.key, temp.value);
            temp = temp.next
        }
        fmt.Print("nil\n");
	}
}
