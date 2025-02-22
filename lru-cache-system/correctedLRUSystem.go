package lrucachesystem

import "errors"

type lruNodeSec struct {
	key   string
	value string
	prev  *lruNodeSec
	next  *lruNodeSec
}

type lruHashMapSec struct {
	cache    map[string]*lruNodeSec
	capacity int
	size     int
	head     *lruNodeSec
	tail     *lruNodeSec
}

func newLruCacheSystem(capacity int) *lruHashMapSec {
	if capacity <= 0 {
		panic("Capacity must be greater than 0")
	}

	return &lruHashMapSec{
		cache:    make(map[string]*lruNodeSec),
		capacity: capacity,
		head:     nil,
		tail:     nil,
		size:     0,
	}
}

func (h *lruHashMapSec) putFunction(key, value string) {
	// checking if the node already exists or nor
	if node, exists := h.cache[key]; exists {
		// Key exists --> move to front
		node.value = value
		h.moveToFront(node)
		return
	}

	//in order to put the new node in the first position in the hashmap
	// we first need to compare the length of the cache with the capacity of the hashmap
	// if the length of the cache is equal to the capacity of the hashmap, then we need to remove the last node of the hashmap
	// and then put the new node in the first position
	if h.size >= h.capacity {
		h.removeLastIndex()
	}
	newNode := &lruNodeSec{key: key, value: value}
	h.addToFront(newNode)
	h.cache[key] = newNode
	h.size++
}

func (h *lruHashMapSec) getFunction(key string) (string, error) {
	if node, exists := h.cache[key]; exists {
		h.moveToFront(node)
		return node.value, nil
	} else {
		return "", errors.New("node specified for the GET operation does not exist")
	}
}

func (h *lruHashMapSec) removeLastIndex() {
	if h.tail == nil {
		return
	}

	delete(h.cache, h.tail.key) // removes the key from the hashmap
	h.tail = h.tail.prev        // moves the tail pointer one position back
	if h.tail == nil {
		h.head = nil // if the tail is also nil, then the head is also nil
	} else {
		h.tail.next = nil // removes the reference of the tail in the previous node
	}

	h.size-- // decreases the size of the hashmap by 1
}

func (h *lruHashMapSec) moveToFront(node *lruNodeSec) {
	if h.head == nil || node == h.head || node == nil {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == h.tail {
		h.tail = node.prev
	}

	node.prev = nil
	node.next = h.head
	if h.head != nil {
		h.head.prev = node
	}
	h.head = node

	if h.tail == nil {
		h.tail = node
	}
}

func (h *lruHashMapSec) addToFront(node *lruNodeSec) {
	if h.head == nil {
		h.head = node
		h.tail = node
		return
	}

	node.next = h.head
	h.head.prev = node
	h.head = node
}
