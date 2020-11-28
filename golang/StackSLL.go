package main

import (
   "fmt"
)

type Node struct { // node using singly linked list
   value int
   link *Node
}

var top *Node = nil // initialize top pointer as zero

func Push(s *Node, value int) { // add item to top
   newNode := new(Node) // create new node in the stack
   newNode.value = value // assign value into new node
   newNode.link = top // new node's next pointer should be nil (which is currently top)
   top = newNode // update top pointer to new node
}

func Pop(s *Node) { // remove the top item
   var temp *Node // temp node for storing pointer location
   temp = top.link // remember the location of next node's pointer
   top = temp // set the next node's pointer to top
}

func Display() {
   var temp *Node // temp is to traverse the stack
   temp = top // set top's pointer to temp
   if(top == nil) { // if top is nil means stack is empty
      fmt.Println("Stack is empty")
   } else { // if stack is not empty
      for temp != nil { // iterate until temp is not nil
         fmt.Println(temp.value) // display current value
         temp = temp.link // go to next location
      }
   }
}

func main() {
   mm := new(Node) // create a stack object
   Push(mm, 7)
   Push(mm, 2)
   Push(mm, 4)
   Push(mm, 8)
   Push(mm, 15)
   Display()
   fmt.Println()

   Pop(mm)
   Pop(mm)
   Display()
}
