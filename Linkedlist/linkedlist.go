package linkedlist

import "fmt"

type LinkedHead struct {
	head   *LinkedNode
	length int
}

type LinkedNode struct {
	data int
	next *LinkedNode
}

var insertedElements = make(map[int]bool)

func (l *LinkedHead) Insert(data int) {
	temp := &LinkedNode{data: data}

	if l.head == nil {
		l.head = temp

	} else {
		temp.next = l.head
		l.head = temp
	}
	l.length++
	insertedElements[data] = true

}

func (l *LinkedHead) Search(data int) {
	if l.head == nil {
		fmt.Println("No Data Found!")
		return
	}
	var length int = 1
	head := l.head
	
	for head.next != nil {
		
		if head.data == data {
			if length==1 {
				fmt.Printf("%d Was The First Node In The Linked List \n", head.data)
			}else{
				fmt.Printf("Data Was Found At %d Number And Data Is:%d", length, head.data)
				return
			}
		}
		head = head.next
		length++
	}
	fmt.Println("No Data Found!")
}

func (l *LinkedHead) Delete(data int) {
	if l.head==nil{
		fmt.Println("No Data Exists!")
	}
	if exists := insertedElements[data]; exists {
		prevNode := l.head

		for prevNode.next != nil {
			if prevNode.next.data == data {
				prevNode.next = prevNode.next.next
				delete(insertedElements,data)
				return
			}
			prevNode = prevNode.next
		}
	} else {
		fmt.Println("No Such Data Exists")
		return
	}
}
func (l *LinkedHead) Show() {
    if l.head == nil {
        fmt.Println("No Data Found!")
        return
    }
    head := l.head
    for head != nil {
        fmt.Printf("%d --> ", head.data)
        head = head.next
    }
    fmt.Println("NULL")
}
	
