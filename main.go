package main

import (
	linkedList "Go_Practice/linked_list/src"
	"fmt"
)

func main() {
	l := linkedList.LinkedList{}

	fmt.Println("\n************* Insert *************")
	l.Insert(12)
	l.Insert(13)
	l.Insert(14)
	l.Insert(15)
	fmt.Println("************* Print *************")
	l.Print()

	fmt.Println("\n************* Search *************")
	fmt.Println("Position of 12 value: ", l.Search(12))
	fmt.Println("Position of 14 value: ", l.Search(14))
	fmt.Println("Position of 15 value: ", l.Search(15))
	fmt.Println("Position of 100 value: ", l.Search(100))

	fmt.Println("\n************* GetAt *************")
	fmt.Println("Get At 1st position: ", l.GetAt(0))
	fmt.Println("Get At 3rd position: ", l.GetAt(2))
	fmt.Println("Get At 4th position: ", l.GetAt(3))
	fmt.Println("Get At -5 position (head is returned): ", l.GetAt(-5))

	// fmt.Println("\n************* InsertAt *************")
	// l.InsertAt(0, 12)
	// l.InsertAt(-1, 13)
	// l.InsertAt(1, 14)
	// l.InsertAt(2, 15)
	// fmt.Println("************* Print *************")
	// l.Print()

	fmt.Println("\n************* DeleteAt *************")
	l.DeleteAt(3)
	fmt.Println("************* Print *************")
	l.Print()

	// fmt.Println("\n************* DeleteVal *************")
	// l.DeleteVal(14)
	// fmt.Println("************* Print *************")
	// l.Print()
}
