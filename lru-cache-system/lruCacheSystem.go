package lrucachesystem

import (
	"errors"
	"fmt"
)

type lruNode struct {
	key   string
	value string
	next  *lruNode
	prev *lruNode
}

type lruHashMap struct {
	bucket []*lruNode
	size    int
}

// initializing a new least recently used system
func newLRUsystem(size int) *lruHashMap {
	return &lruHashMap{
		bucket: make([]*lruNode, size),
		size:    size,
	}
}

// this LRU system will be different from that of an ordinary hashmaps
// there will be only one key-value pair corresponding to an index in the array
// but in case of collision, there will be no key-value pair put next to the existing node
// simply, the whole system will be checked, if not at capacity, then it will be put in first
// but, if at capacity, it will be put in first and the last one will be removed


// some edge cases in case of put function
// if the first index is nil, then the key-value pair will be put in first index only
// if the first index is not nil, then the existing key-value pairs in those indices will be pushed below, may leading to the deletion of a pair, then the new pair will be put in the first index
// if the key of the new key-value matches with an existing key of a key-value pair, then an error will be shown of duplicate key present in there

func (h *lruHashMap) PutFunction(key, value string) (string, error) {
	newNode := &lruNode{key: key, value: value};
	index := 0;

	var isKeyPresent = h.IsKeyPresent(key);
	if isKeyPresent {
		return "", errors.New("duplicate key already present");
	}

	// head position is empty, then put it in the first index
	if h.bucket[index] == nil {
		h.bucket[index] = newNode
		return "Node put in the first position, as head was empty", nil;
	} else {
		// if the head is not empty, then start from the next index till the last one
		for index = 1; index < h.size; index++ {
			temp := h.bucket[index]
			prevTemp := h.bucket[index]
			if temp == nil {
				// nodes are moved one after their respective positions
				prevTemp.next = temp.next;
				temp = prevTemp;

				// the first position becomes empty, so the new node is put here
				h.bucket[0] = newNode
				return "Node put in head, some nodes were have to be moved, but it's alright", nil;
			} else {
				temp = temp.next
				// tail situation
				if temp.next == nil {
					h.RemoveLastIndex(newNode)
					return "last node was removed, the subsequent nodes were moved and the new node was put in the first position", nil
				}
				// else continue the loop
			}
		}
	}

	return "", errors.New("the new node was not put in the hash map due to some problems");
}

// checking whether the key is already present
func (h *lruHashMap) IsKeyPresent(key string) bool {
	head := h.bucket[0];
	temp := head;

	for temp != nil {
		if temp.key == key {
			return true
		} else {
			temp = temp.next
		}
	}

	return false
}

// Get functionality done
func (h *lruHashMap) GetFunction(key string) (string, error) {
	index := 0;
	temp := h.bucket[index];
	for temp != nil {
		if temp.key == key {
			return temp.value, nil
		} else {
			temp = temp.next
		}
	}

	return "", errors.New("given key not found");
}


// Remove functionality done
func (h *lruHashMap) RemoveLastIndex(newNode *lruNode) {
	// if the last index is not nil/empty, then the whole hashmap is at capacity
	index := h.size - 1;
	head := h.bucket[0];
	if h.bucket[index] == nil {
		fmt.Print("this last index is empty")
	} else {
		// the node at last index is removed
		last := h.bucket[index];
		secLast := h.bucket[index-1]
		secLast.next = last.next;
		last.prev = nil;

		// and the new node is put in the first position
		// replacing the head
		head.prev = newNode
		newNode.next = head
		head = newNode
	}
}