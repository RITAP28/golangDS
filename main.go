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
	ll.DisplayLinkedList();

	ll.InsertAtBeginning(2);
	ll.DisplayLinkedList();

	ll.InsertAtBeginning(3);
	ll.DisplayLinkedList();

	ll.InsertAtBeginning(4);
	ll.DisplayLinkedList();

	ll.InsertAtEnding(100);
	ll.DisplayLinkedList();

	ll.InsertAtEnding(200);
	ll.DisplayLinkedList();
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