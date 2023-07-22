package main

import "fmt"

// ArraySize is the size of hash table array
const ArraySize = 7

// HashTable will hold an array
type HashTable struct{
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the Hash table
type bucket struct{
	head *bucketNode
}

// bucketNode structure
type bucketNode struct{
	key string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string){
	index := hash(key)

	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool{
	index := hash(key)

	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string){
	index := hash(key)

	h.array[index].delete(key)
}

// insert will take in a key, create a node with that key and insert the node on the bucket
func (b *bucket) insert(k string){
	if b.search(k){
		fmt.Println(k, "already exists")
		return
	}

	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// search
func (b *bucket) search(k string) bool{
	currNode := b.head

	for currNode != nil{
		if currNode.key == k{
			return true
		}

		currNode = currNode.next
	}

	return false
}

// delete will take in a key and delete it from the bucket
func (b *bucket) delete(k string) {
	if !b.search(k){
		fmt.Println(k, "doesn't exist")
		return
	}

	if b.head.key == k{
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil{
		if prevNode.next.key == k{
			prevNode.next = prevNode.next.next
		}
		
		prevNode = prevNode.next
	}

}

// hash
func hash(key string) int{
	sum := 0

	for _, c := range key{
		sum += int(c)
	}

	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable{
	result := &HashTable{}

	for i := range result.array{
		result.array[i] = &bucket{}
	}

	return result
}

func main(){
	hashTable := Init()

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _,v := range list{
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
	fmt.Println(hashTable.Search("STAN"))
	fmt.Println(hashTable.Search("KENNY"))
}