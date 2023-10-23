package benchmarkBinarySearchTree

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

var headNode BSTNode
var dataSize int

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

func TestBST(d_Size int) {
	dataSize = d_Size

	var arr_randomData []int = []int{}
	for i := 1; i < dataSize; i++ {
		arr_randomData = append(arr_randomData, i)
	}

	rand.Shuffle(len(arr_randomData), func(i, j int) {
		arr_randomData[i], arr_randomData[j] = arr_randomData[j], arr_randomData[i]
	})

	for _, elem := range arr_randomData {
		var insertTime = time.Now().UnixNano()
		insertData(elem)
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
		accessData(headNode)
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
		searchData(headNode, rand.Intn(dataSize-0)+0)
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

	fmt.Println("|=========== Completed Binary Search Tree ===========|")
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
	fmt.Println("|=============================================|")
}

func accessData(currentNode BSTNode) {
	if currentNode.left != nil {
		accessData(*currentNode.left)
	}

	if currentNode.right != nil {
		accessData(*currentNode.right)
	}
}

func searchData(currentNode BSTNode, value int) (*BSTNode, error) {
	if currentNode.data == value {
		return &currentNode, nil
	}

	if value < currentNode.data && currentNode.left != nil {
		return searchData(*currentNode.left, value)
	}

	if value > currentNode.data && currentNode.right != nil {
		return searchData(*currentNode.right, value)
	}

	return nil, errors.New("valueNotFound")
}

func insertData(value int) {
	if headNode.data == 0 {
		headNode = BSTNode{
			data:  value,
			right: nil,
			left:  nil,
		}
		return
	}

	var iter int = 0
	var cursor *BSTNode = &headNode

	for iter < dataSize {
		if cursor.data > value {
			if cursor.left != nil {
				cursor = cursor.left
				continue
			} else {
				cursor.left = &BSTNode{
					data:  value,
					right: nil,
					left:  nil,
				}

				return
			}
		}

		if cursor.data < value {
			if cursor.right != nil {
				cursor = cursor.right
				continue
			} else {
				cursor.right = &BSTNode{
					data:  value,
					right: nil,
					left:  nil,
				}

				return
			}
		}

		iter++
	}
}

func deleteData(value *BSTNode) error {
	return nil
}

//TODO sorting
