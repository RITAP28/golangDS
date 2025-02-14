package datastructures

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

	return ""
}
