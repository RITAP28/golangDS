package main

import (
	// "errors"
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	lrucachesystem "test-lru-cache-system/lru-cache-system"
)

const (
	PORT       = ":6379"
	CACHE_SIZE = 5
)

var cache = lrucachesystem.NewLruCacheSystem(CACHE_SIZE)

func handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// Read Client Input
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Client disconnected: ", err)
			return
		}

		// process command
		command = strings.TrimSpace(command)
		parts := strings.Split(command, " ")

		switch strings.ToUpper(parts[0]) {
		case "SET":
			if len(parts) < 3 {
				conn.Write([]byte("ERROR: SET requires key and value\n"))
				continue
			}
			key, value := parts[1], parts[2]
			cache.PutFunction(key, value)
			conn.Write([]byte("OK\n"))

		case "GET":
			if len(parts) < 2 {
				conn.Write([]byte("ERROR: GET requires a key\n"))
				continue
			}
			key := parts[1]
			if value, found := cache.GetFunction(key); found {
				conn.Write([]byte(value + "\n"))
			} else {
				conn.Write([]byte("(nil)\n"))
			}

		case "EXIT":
			conn.Write([]byte("Goodbye!\n"))
			return

		default:
			conn.Write([]byte("ERROR: Unknown Command\n"))
		}
	}
}

func main() {
	// fmt.Println("I am a boy");
	// firstName, lastName := getNames();
	// fmt.Printf("The first name is %v and last name is %v\n", firstName, lastName);

	// sum := 20;
	// fmt.Println(split(sum));

	// ll := datastructures.LinkedList{}

	// ll.InsertAtBeginning(1)
	// ll.InsertAtBeginning(2)
	// ll.InsertAtBeginning(3)
	// ll.InsertAtBeginning(4)
	// ll.InsertAtEnding(100)
	// ll.InsertAtEnding(200)
	// ll.DisplayLinkedList()

	// dll := datastructures.DoublyLinkedList{}
	// dll.InsertAtBeginning(1)
	// dll.InsertAtBeginning(2)
	// dll.ForwardTraversing()

	// hashMap := datastructures.NewHashMap(10)
	// fmt.Println()
	// fmt.Println("Hashmap being dispayed here ----> ")
	// hashMap.PutFunction("apple", "10")
	// hashMap.PutFunction("banana", "20")
	// hashMap.PutFunction("guava", "30")
	// hashMap.PutFunction("table", "69")
	// hashMap.Display()
	// fmt.Println()
	// fmt.Print(hashMap.GetFunction("table"))
	// fmt.Println()
	// fmt.Println("Hashmap being displayed after deletion of a particular key and string ---->")
	// hashMap.DeleteFunction("apple")
	// hashMap.Display()

	fmt.Println("Starting LRU Cache Server on", PORT)

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			continue
		}

		go handleClient(conn)
	}
}

// func getNames() (string, string) {
// 	return "John", "Doe"
// }

// func split(sum int) (x, y int, err error) {
// 	if sum < 20 {
// 		err = errors.New("sum cannot be less than 20");
// 		return 0, 0, err
// 	}
// 	x = sum * 4 / 10;
// 	y = sum - x;
// 	return x, y, nil;
// }
