package main

import (
	"fmt"

	benchmarkArray "github.com/RooopDooop/DataStructuresBenchmark/Array"
	benchmarkBinarySearchTree "github.com/RooopDooop/DataStructuresBenchmark/BinarySearchTree"
	benchmarkLinkedList "github.com/RooopDooop/DataStructuresBenchmark/LinkedList"
)

//Data structures TODO
//Array DONE (Except Delete)
//Linked List (Except Delete)
//BST TODO
//Graph TODO
//Stack TODO
//Matrix
//Heap
//Queue
//Hashing

func main() {
	fmt.Println("Starting Datastructure Benchmarks...")
	benchmarkArray.TestArray(100000)
	benchmarkLinkedList.TestLinkedList(100000)
	benchmarkBinarySearchTree.TestBST(100000)
	fmt.Println("Finished Datastructure Benchmarks...")
}
