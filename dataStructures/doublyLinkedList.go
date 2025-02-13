package datastructures

import "fmt"

type doubleNode struct {
	value int
	next *doubleNode
	prev *doubleNode
}

type DoublyLinkedList struct {
	head *doubleNode
	tail *doubleNode
}

// inserting the node at the beginning of the list
func (dll *DoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &doubleNode{value: data};
	if dll.head == nil {
		dll.head = newNode;
		dll.tail = newNode;
		return;
	}
	newNode.next = dll.head;
	dll.head.prev = newNode;
	dll.head = newNode;
}

// forward traversing through the list
func (dll *DoublyLinkedList) ForwardTraversing() {
	temp := dll.head;
	fmt.Println("started forward traversing the dll");
	for ; temp != nil; {
		fmt.Printf("%d --> ", temp.value);
		temp = temp.next;
	}
	fmt.Print("nil");
	fmt.Println();
	fmt.Println("completed forward traversing the dll");
}

