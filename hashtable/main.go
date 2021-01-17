// https://www.youtube.com/watch?v=zLnJcAt1aKs
// https://github.com/lee-junmin/tutorial-source/blob/master/data-structure-golang/hash-table/main.go
package main

import (
	"fmt"
)

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable structure will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure is a linked list in each slot of the hash table array
type bucket struct {
	head *bucketNode
}

// bucketNode structure is a linked list node that holds the key and address of
// the next node
type bucketNode struct {
	key  string
	next *bucketNode
}

// HashTable methods
// Insert will take in a key and add it to the hash tabl array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash
// table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Print will print the buckets in each slot of the hash table array
func (h *HashTable) Print() {
	for i := 0; i < len(h.array); i++ {
		h.array[i].print(i)
	}
}

// bucket methods
// insert will take in a key, create a node with the key, and insert the node
// in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("%s already exists\n", k)
	}
}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head

	// loop until currentNode is empty (nil)
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}

		// move to next node
		currentNode = currentNode.next
	}

	// if we reach this point, the end of the linked list was reached without a
	// match
	return false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	// edge case: if match is the head node
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	// remove node by linking the previous node to the node after
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode = previousNode.next.next
			return
		}

		// move to the next node
		previousNode = previousNode.next
	}
}

// print will print all the bucketNode keys in the bucket
func (b *bucket) print(slot int) {

	fmt.Printf("HashTable slot %v: ", slot)

	// print emtpy if there are no bucketNodes
	if b.head == nil {
		fmt.Printf("bucket is empty")
	}

	arrow := ""
	currentNode := b.head
	for currentNode != nil {
		if currentNode.next == nil {
			arrow = ""
		} else {
			arrow = "=> "
		}
		fmt.Printf("%s %s", currentNode.key, arrow)
		currentNode = currentNode.next
	}
	fmt.Println("")
}

// hash
func hash(key string) int {
	// sum the ASCII values for each letter in the key string
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	// return the remainder of sum divided by ArraySize
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}

	// create a bucket for each slot
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	// Testing statements
	// testHashTable := Init()
	// fmt.Println(testHashTable)
	// fmt.Println(hash("RANDY"))

	// testBucket := &bucket{}
	// testBucket.insert("RANDY")
	// testBucket.delete("RANDY")
	// fmt.Println(testBucket.head.key)

	// fmt.Println(testBucket.search("RANDY"))
	// fmt.Println(testBucket.search("ERIC"))

	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
		"MARK",
	}

	fmt.Printf("Adding the following keys to the hash table: ")
	fmt.Println(list)
	fmt.Println("")

	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println("The hash table now contains these buckets.")
	hashTable.Print()

	fmt.Println("Removing ERIC")
	hashTable.Delete("ERIC")

	fmt.Println("")
	fmt.Println("The hash table now contains these buckets.")
	hashTable.Print()

}
