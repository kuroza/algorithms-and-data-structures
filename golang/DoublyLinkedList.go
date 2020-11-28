package main

import (
   "fmt"
)

type List struct {
   head *Node
   tail *Node
   count int
}

type Node struct {
   value int
   next *Node
   prev *Node
}

// Create empty DLL
func (list *List) MakeList() {
   newNode := new(Node) // create an empty new node
   list.head = newNode
   list.tail = newNode
}

func (list *List) AddAtBeginning(value int) { // DLL
   newNode := new(Node) // create a new node
   if list.count == 0 { // if list is empty, point head and tail to new node
      list.head = newNode
      list.tail = newNode
   } else {
      list.head.prev = newNode // set the current head node's previous pointer to new node
      newNode.next = list.head // set the current head node to the new node's next pointer
      list.head = newNode // update the head to point to the new node
   }
   list.count++
}

func (list *List) IsEmpty() bool {
   return list.count == 0
}

func (list *List) Size() int {
   return list.count
}

func main() {
   SLL := new(List)
   SLL.MakeList()
   SLL.AddAtBeginning(2)
   SLL.AddAtBeginning(27)

   if SLL.IsEmpty() == true {
      fmt.Println("The list is empty.")
   } else {
      fmt.Println("The list is NOT empty.")
   }

   fmt.Println("Size of list: ", SLL.Size())
}
