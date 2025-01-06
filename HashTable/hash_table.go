package main

import (
	"fmt"
)

// HashTable Structure
const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket structure
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].Delete(key)
}

// Insert into bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// Show hash table contents
func (h *HashTable) Show() {
	for i := 0; i < ArraySize; i++ {
		head := h.array[i].head
		if head == nil {
			fmt.Printf("Bucket %d: (empty)\n", i)
		} else {
			fmt.Printf("Bucket %d:\n", i)
			for head != nil {
				PrintWithSeparator("value:", head.key)
				head = head.next
			}
		}
	}
}

// Print with separator
func PrintWithSeparator(output1, output2 string) {
	fmt.Printf("  %-10s | %s\n", output1, output2)
}

// Search in bucket
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

// Delete from bucket
func (b *bucket) Delete(k string) {
	if b.head == nil {
		return
	}
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

// Hash Function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Initialize hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// Main function
func main() {
	hashTable := Init()
	list := []string{
		"Guatam",
		"Krishna",
		"Radha",
		"Balram",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	// Display hash table contents
	hashTable.Show()
}
