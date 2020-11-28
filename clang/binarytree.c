#include <stdio.h>
#include <stdlib.h>

struct node
{
   int data;
   struct node *left, *right;
}

struct node *create()
{
   int x;
   struct node *newnode;
   newnode = (struct node *) malloc(sizeof(struct node)) // a new node will be created
   printf("Enter data (-1 for no node): ");
   scanf("%d", &x); // input data
   if(x == -1) // if data is -1, return back to previous node
   {
      return 0;
   }
   newnode->data = x;
   printf("Enter left child of %d", x);
   newnode->left = create(); // recursion for left child
   printf("Enter right child of %d", x); // once the left child have no more nodes, it will return and ask for the right child
   newnode->right = create(); // recursion for right child
   return newnode; // this function returns the address of the root node
}

void main()
{
   struct node *root;
   root = 0;
   root = create(); // where root will be stored
}
