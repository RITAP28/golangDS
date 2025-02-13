package datastructures

import "fmt"

// Node represents a node in a linked list
type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{value: data};
	if ll.head == nil {
		ll.head = newNode;
		ll.tail = newNode;
		return;
	}
	newNode.next = ll.head;
	ll.head = newNode;
}

func (ll *LinkedList) InsertAtEnding(data int) {
	newNode := &Node{value: data};
	if ll.head == nil {
		ll.tail = newNode;
		ll.head = newNode;
		return;
	}

	ll.tail.next = newNode;
	ll.tail = newNode;
}

func (ll *LinkedList) DisplayLinkedList() {
	if ll.head == nil {
		fmt.Println("The linked list is empty");
		return;
	};
	current := ll.head;
	fmt.Println("traversing through the linked list: ")
	for ; current != nil; {
		fmt.Printf("%d --> ", current.value);
		current = current.next;
	};
	fmt.Print("nil");
	fmt.Println();
	fmt.Println("completed traversing the linked list");
}