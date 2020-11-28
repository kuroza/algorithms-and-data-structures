package main

import (
   "fmt"
)

type List struct { // SLL has no tail
   head *Node // head pointer of the list
   count int // stores the node counts in the list
}

type Node struct { // SLL has no prev
   value int // stores the value of a node
   next *Node // stores the next node's pointer
}

func (list *List) InsertBeginning(value int) { // SLL
   list.head = &Node{value, list.head} // the newly created node's next pointer will contain the previous head's pointer
   // then head will point to the new node
   list.count++ // increment the amount of nodes in the list
}


func (list *List) InsertEnd(value int) { // insert a value at the end of a list
   temp := list.head // store head pointer to temp
   newNode := &Node{value, nil} // create new node, nil because it's at the end
   list.count++ // increment count
   if temp == nil { // if the list is empty
      list.head = newNode // set the head to the new node
      return // finish
   }
   for temp.next != nil { // iterate until the end (until the pointer is not nil)
      temp = temp.next // traverse, go to next node
   }
   temp.next = newNode // set a pointer to the newly created node at the tail node in the list
}

func (list *List) PrintList() { // print all values in the list
   temp := list.head // temp pointer is for traversing the list, starts from the head
   for temp != nil {
      fmt.Print(temp.value, " ") // print value
      temp = temp.next // go to next node
   }
   fmt.Println()
}

func (list *List) FirstItem() int { // prints the first item in the list
   return list.head.value // returns the first node's value
}

func (list *List) LastItem() int { // prints the last item in the list
   temp := list.head // temp is for traversing the list
   for temp.next != nil {
      temp = temp.next // store the last node's pointer
   }
   return temp.value // returns the last node's value
}

func (list *List) RemoveFirst() {
   list.head = list.head.next // changes the head pointer to the next node's pointer, hence removing the first node from the list
   list.count-- // decrement the list count
}

func (list *List) RemoveLast() { // remove last needs 2 pointers
   temp := list.head // to traverse the list, initially temp and prevNode = head
   prevNode := list.head // prevNode is to change the second last pointer to nil
   if temp.next == nil { // if there's only 1 node, delete that node by head = 0
      list.head = nil
   }
   for temp.next != nil {
      prevNode = temp // store current pointer to prevNode
      temp = temp.next // update temp's to next node
   }
   prevNode.next = nil // update the prevNode to nil, i.e. the end is removed from the list
   list.count-- // decrement the count once a node's been removed
}

func (list *List) Size() int {
   return list.count // get the list size that is stored in count
}

func (list *List) Search(value int) bool {
   temp := list.head
   for temp != nil { // traverse
      if temp.value == value { // if value is found
         return true
      }
      temp = temp.next // go to next node
   }
   return false // return false if value is not found
}

func (list *List) Sum() int {
   temp := list.head // temp is a pointer for traversing the list
   var sum int // declare sum
   for temp != nil { // traverse the list
      sum = sum + temp.value // store values into sum
      temp = temp.next
   }
   return sum
}

func main() {
   SLL := new(List) // create a new list object
   SLL.InsertBeginning(5)
   SLL.InsertBeginning(14)
   SLL.InsertEnd(2000)
   SLL.InsertEnd(62)
   SLL.InsertBeginning(40)

   fmt.Println("First item: ", SLL.FirstItem())
   fmt.Println("Last item: ", SLL.LastItem())

   fmt.Print("List: ")
   SLL.PrintList()
   fmt.Println("Length of the list: ", SLL.Size())
   fmt.Println("Sum: ", SLL.Sum())

   SLL.RemoveFirst()
   fmt.Print("List after removing first item: ")
   SLL.PrintList()

   SLL.RemoveLast()
   fmt.Print("List after removing last item: ")
   SLL.PrintList()

   fmt.Println("Length of the list: ", SLL.Size())
   fmt.Println("Sum: ", SLL.Sum())
}
