package main

import (
   "fmt"
)

type Node struct { // constructor for a node
   value int
   left *Node
   right *Node
}

func ArrayToBST(arr []int) *Node {
   var root *Node
   i := 0
   for i < len(arr) {
      root = Insert(root, arr[i])
      i++
   }
   return root
}

func Insert(root *Node, value int) *Node {
   if root == nil { // if the BST is empty
      root = &Node{value, nil, nil} // create new node
   } else if value < root.value { // if value is less then go left
      root.left = Insert(root.left, value) // recursion: go left.. go left.. go left.. insert value
   } else { // else go right
      root.right = Insert(root.right, value) // recursion: go right.. go right.. go right.. insert value
   }
   return root // return the root node
}

func PrintInOrder(root *Node) { // Printing: Left Root Right, sorted
   if root == nil {
      return
   }
   PrintInOrder(root.left) // recursion left
   fmt.Print(root.value, " ") // print root
   PrintInOrder(root.right) // recursion right
}

func PrintPreOrder(root *Node) { // Printing: Root Left Right
   if root == nil {
      return
   }
   fmt.Print(root.value, " ") // print root
   PrintPreOrder(root.left) // recursion left
   PrintPreOrder(root.right) // recursion right
}

func PrintPostOrder(root *Node) { // Printing: Left Right Root
   if root == nil {
      return
   }
   PrintPostOrder(root.left) // recursion left
   PrintPostOrder(root.right) // recursion right
   fmt.Print(root.value, " ") // print root
}

func Search(root *Node, value int) *Node {
   temp := root // temp is a pointer to traverse
   for temp != nil {
      if temp.value == value { // if value is found
         return temp // return a pointer
      } else if value < temp.value { // if the value is less
         temp = temp.left // traverse left
      } else {
         temp = temp.right // else traverse right
      }
   }
   return nil // if not found, return nil
}

func SearchSmallest(root *Node) int { // to find smallest, go all the way to the left most of tree
   temp := root // temp is for traversing the list
   if temp == nil { // if the tree is empty
      fmt.Println("EmptyTreeException")
      return 0
   }
   for temp.left != nil { // traverse left until end, because BST small is on left side
      temp = temp.left
   }
   return temp.value // return value
} // O(n)

func RemoveSmallest(root *Node) {
   temp := root
   prevNode := root
   for temp.left != nil {
      prevNode = temp
      temp = temp.left
   }
   prevNode.left = nil
} // O(n)

func SearchLargest(root *Node) int { // to find largest, go all the way to the right most
   temp := root
   if temp == nil {
      fmt.Println("EmptyTreeException")
      return 0
   }
   for temp.right != nil {
      temp = temp.right
   }
   return temp.value
} // O(n)

func TreeDepth(root *Node) int {
   if root == nil {
      return 0
   }
   lDepth := TreeDepth(root.left)
   rDepth := TreeDepth(root.right)

   if lDepth > rDepth {
      return lDepth + 1
   }
   return rDepth + 1
}

func Sum(root *Node) int {
   var sum, leftSum, rightSum int
   if root == nil {
      return 0
   }
   rightSum = Sum(root.right)
   leftSum = Sum(root.left)
   sum = rightSum + leftSum + root.value
   return sum
}

func main() {
   var tree *Node //create a tree object
   arr := []int{4, 10, 8, 6, 1, 3, 9}
   tree = ArrayToBST(arr)
   PrintInOrder(tree)
   fmt.Println()

   tree = Insert(tree, 5)
   tree = Insert(tree, 7)
   tree = Insert(tree, 2)
   fmt.Print("In-order traversal: ")
   PrintInOrder(tree)
   fmt.Println()

   fmt.Print("Pre-order traversal: ")
   PrintPreOrder(tree)
   fmt.Println()

   fmt.Println("Pointer to search value: ", Search(tree, 6)) // prints a pointer if found
   fmt.Println("Smallest value: ", SearchSmallest(tree)) // prints the smallest integer
   RemoveSmallest(tree)
   PrintInOrder(tree)
   fmt.Println(Sum(tree))

}
