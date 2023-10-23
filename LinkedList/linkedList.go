package benchmarkLinkedList

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	data int
	next *Node
}

var headNode Node

//var iterationCounts int = 500000

var insertAvg []int64
var insertAvgTotal int64 = 0
var insertHighest int64
var insertLowest int64 = 9223372036854775807

var accessAvg []int64
var accessAvgTotal int64 = 0
var accessHighest int64
var accessLowest int64 = 9223372036854775807

var searchingAvg []int64
var searchingAvgTotal int64 = 0
var searchingHighest int64
var searchingLowest int64 = 9223372036854775807

var deleteAvg []int64
var deleteHighest int64
var deleteLowest int64

func TestLinkedList(dataSize int) {
	//for i := 1; i < 500000; i++ {

	//Temporary for development
	for i := 1; i < dataSize; i++ {
		var insertTime = time.Now().UnixNano()
		insertData(i)
		var insertTimeEnd = time.Now().UnixNano()

		insertAvg = append(insertAvg, (insertTimeEnd - insertTime))
		insertAvgTotal = insertAvgTotal + int64((insertTimeEnd - insertTime))

		if (insertTimeEnd - insertTime) > insertHighest {
			insertHighest = (insertTimeEnd - insertTime)
			continue
		}

		if (insertTimeEnd - insertTime) < insertLowest {
			insertLowest = (insertTimeEnd - insertTime)
			continue
		}
	}

	for i := 1; i < dataSize; i++ {
		var accessTime = time.Now().UnixNano()
		accessData()
		var accessTimeEnd = time.Now().UnixNano()

		accessAvg = append(accessAvg, (accessTimeEnd - accessTime))
		accessAvgTotal = accessAvgTotal + int64((accessTimeEnd - accessTime))

		if (accessTimeEnd - accessTime) > accessHighest {
			accessHighest = (accessTimeEnd - accessTime)
			continue
		}

		if (accessTimeEnd - accessTime) < accessLowest {
			accessLowest = (accessTimeEnd - accessTime)
			continue
		}
	}

	for i := 1; i < dataSize; i++ {
		var searchingTime = time.Now().UnixNano()
		searchData(rand.Intn(dataSize-0) + 0)
		var searchingTimeEnd = time.Now().UnixNano()

		searchingAvg = append(searchingAvg, (searchingTimeEnd - searchingTime))
		searchingAvgTotal = searchingAvgTotal + int64((searchingTimeEnd - searchingTime))

		if (searchingTimeEnd - searchingTime) > searchingHighest {
			searchingHighest = (searchingTimeEnd - searchingTime)
			continue
		}

		if (searchingTimeEnd - searchingTime) < searchingLowest {
			searchingLowest = (searchingTimeEnd - searchingTime)
			continue
		}
	}

	/*var searchingTime = time.Now().UnixNano()
	deleteTarget, delErr := searchData(rand.Intn(9999-0) + 0)
	var searchingTimeEnd = time.Now().UnixNano()
	if delErr != nil {
		log.Fatalln(delErr)
	}*/

	/*var deleteTime = time.Now().UnixNano()
	deleteData(deleteTarget)
	var deleteTimeEnd = time.Now().UnixNano()*/

	fmt.Println("|=========== Completed Linked List ===========|")
	fmt.Printf("	Insert: \n")
	fmt.Printf("		- Avg: %d \n", (insertAvgTotal / int64(len(insertAvg))))
	fmt.Printf("		- Low: %d \n", insertLowest)
	fmt.Printf("		- High: %d \n", insertHighest)

	fmt.Printf("	Access: \n")
	fmt.Printf("		- Avg: %d \n", (accessAvgTotal / int64(len(accessAvg))))
	fmt.Printf("		- Low: %d \n", accessLowest)
	fmt.Printf("		- High: %d \n", accessHighest)

	fmt.Printf("	Searching: \n")
	fmt.Printf("		- Avg: %d \n", (searchingAvgTotal / int64(len(searchingAvg))))
	fmt.Printf("		- Low: %d \n", searchingLowest)
	fmt.Printf("		- High: %d \n", searchingHighest)

	//fmt.Printf(" - Searching: %d\n", searchingTimeEnd-searchingTime)
	//fmt.Printf(" - Deleting: %d\n", deleteTimeEnd-deleteTime)
	fmt.Println("|=============================================|")
}

func accessData() {
	var cursor = &headNode
	for {
		if cursor.next == nil {
			return
		}

		cursor = cursor.next
	}
}

func searchData(value int) (*Node, error) {
	var cursor = &headNode
	for {
		if cursor.next == nil {
			return nil, errors.New("dataNotFound")
		}

		if cursor.data == value {
			return cursor, nil
		}

		cursor = cursor.next
	}
}

func insertData(value int) {
	var newData = Node{data: value, next: nil}

	var cursor = &headNode
	for {
		if cursor.next == nil {
			cursor.next = &newData
			return
		}

		cursor = cursor.next
	}
}

func deleteData(value *Node) error {
	var cursor = &headNode
	for {
		if cursor.next == nil {
			return errors.New("dataNotFound")
		}

		if cursor.next == value {
			cursor.next = cursor.next.next
			return nil
		}

		cursor = cursor.next
	}
}

//TODO sorting
