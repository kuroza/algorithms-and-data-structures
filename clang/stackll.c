#include <stdio.h>
#include <stdlib.h>

struct node
{
   int data;
   struct node *link;
}

struct node *top = 0;

void push(int x)
{
   struct node *newnode;
   newnode = (struct node *) malloc(sizeof(struct node));
   newnode->data = x;
   newnode->next = top;
   top = newnode;
}

void display()
{
   struct node *temp;
   temp = top;
   if(top == 0)
   {
      printf("Stack is empty");
   }
   else
   {
      while(temp != 0)
      {
         printf("%d", temp->data);
         temp = temp->link;
      }
   }
}
