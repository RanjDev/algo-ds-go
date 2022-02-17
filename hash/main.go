package main

import "fmt"

const ArraySize = 7

// Hashtable structure
type Hashtable struct {
	array [ArraySize]*bucket
}

// Bucket structure, linkedlist
type bucket struct {
	head *bucketNode
}

// BuckeNode structure, node
type bucketNode struct {
	key  string
	next *bucketNode
}

// hashtable operations
// insert
func (h *Hashtable) insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// search
func (h *Hashtable) search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

// delete
func (h *Hashtable) delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket operations
// insert
func (b *bucket) insert(k string) {
	if !b.search(k) {

		newNode := &bucketNode{key: k}

		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Already Exists")
	}
}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head

	for prevNode.next != nil {

		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
			return
		}

		prevNode = prevNode.next

	}

}

// hash
func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init
// create a bucket in each slot of the hashtable
func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	fmt.Println("Hashtable")

	testHash := Init()

	testHash.insert("Randy")
	testHash.insert("Eric")
	fmt.Println(testHash.search("Eric"))
	fmt.Println(testHash.search("bob"))

	testHash.delete("Randy")

	fmt.Println(testHash)

}
