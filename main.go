package main

import (
	"errors"
	// "fmt"
	datastructures "test-lru-cache-system/dataStructures"
)

func main() {
	// fmt.Println("I am a boy");
	// firstName, lastName := getNames();
	// fmt.Printf("The first name is %v and last name is %v\n", firstName, lastName);

	// sum := 20;
	// fmt.Println(split(sum));

	ll := datastructures.LinkedList{};

	ll.InsertAtBeginning(1);
	ll.InsertAtBeginning(2);
	ll.InsertAtBeginning(3);
	ll.InsertAtBeginning(4);
	ll.InsertAtEnding(100);
	ll.InsertAtEnding(200);
	ll.DisplayLinkedList();

	dll := datastructures.DoublyLinkedList{};
	dll.InsertAtBeginning(1);
	dll.InsertAtBeginning(2);
	dll.ForwardTraversing();
}

func getNames() (string, string) {
	return "John", "Doe"
}

func split(sum int) (x, y int, err error) {
	if sum < 20 {
		err = errors.New("sum cannot be less than 20");
		return 0, 0, err
	}
	x = sum * 4 / 10;
	y = sum - x;
	return x, y, nil;
}