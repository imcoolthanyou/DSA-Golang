package linkedlist

import "fmt"

type DoubleLinkedListNode struct {
	PrevNode *DoubleLinkedListNode
	Data     int
	NextNode *DoubleLinkedListNode
}

type DoubleLinkedList struct {
	Head *DoubleLinkedListNode
}

var values= make(map[int]bool)

func (d DoubleLinkedList) Insert(val int) {
	temp := &DoubleLinkedListNode{Data: val}

	if d.Head == nil {
		d.Head=temp

	}else{
		temp.NextNode=d.Head
		d.Head.PrevNode=temp
		d.Head=temp
	}
	values[val]=true
}

func (d DoubleLinkedList) Search(val int){
	if d.Head==nil{
		fmt.Print("No Data Available!")
		return
	}
	

	head:=d.Head
	position:=0

	for head.NextNode!=nil{
		if head.Data==val{
			fmt.Printf("Data is:%d,Present At The %d Number LinkedList!",val,position+1)
	
			return
		}
		
		head=head.NextNode
		position++
	}
	fmt.Println("No Dat Found!")

}

func (d *DoubleLinkedList) Delete(val int) {
    if d.Head == nil {
        fmt.Println("No Data Available!")
        return
    }

    if d.Head.Data == val {
        d.Head = d.Head.NextNode
        if d.Head != nil {
            d.Head.PrevNode = nil
        }
        return
    }

    prevNode := d.Head
    deleted := false

    for prevNode.NextNode != nil {
        if prevNode.NextNode.Data == val {
            prevNode.NextNode = prevNode.NextNode.NextNode
            if prevNode.NextNode != nil {
                prevNode.NextNode.PrevNode = prevNode
            }
            deleted = true
            break
        }
        prevNode = prevNode.NextNode
    }

    if !deleted {
        fmt.Println("No Such Value To Delete!")
    }
}
