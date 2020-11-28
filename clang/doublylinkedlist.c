#include <stdio.h>
#include <stdlib.h>

struct node
{
   int data;
   struct node *next;
   struct node *prev;
};

struct node *head, *tail;

void createDLL()
{
   struct node *newnode, head = NULL; // initialize head as NULL
   int choice = 1;
   while(choice)
   {
      newnode = (struct node *) malloc(sizeof(struct node)); // create a newnode
      printf("Enter data: ");
      scanf("%d", &newnode->data); // get data from user
      newnode->prev = NULL, newnode->next = NULL; // initialize NULL to newnode's pointers
      if(head == NULL)
      {
         head = tail = newnode; // for first node
      }
      else
      {
         newnode->prev = tail; // set newnode's previous pointer
         tail->next = newnode; // set newnode's next pointer
         tail = newnode; // set tail to newnode
      }
      printf("Do you want to create another node? ");
      scanf("%d", &choice);
   }
}

void insertAtBeg() // PUSH
{
   struct node *newnode;
   newnode = (struct node*) malloc(sizeof(struct node));
   printf("Enter data: ");
   scanf("%d", &newnode->data);
   newnode->prev = NULL, newnode->next = NULL;
   head->prev = newnode; // update current node's pointer
   newnode->next = head; // or the newode's pointer first
   head = newnode; // then change the head/tail pointer
}

void insertAtEnd()
{
   struct node *newnode;
   newnode = (struct node*) malloc(sizeof(struct node));
   printf("Enter data: ");
   scanf("%d", &newnode->data);
   newnode->prev = NULL, newnode->next = NULL;
   tail->next = newnode; // if there's tail
   newnode->prev = tail;
   tail = newnode;
}

void insertAtPosition()
{
   int pos, i=1;
   printf("Enter position: ");
   scanf("%d", &pos);
   struct node *newnode, *temp;
   temp = head;
   newnode = (struct node*) malloc(sizeof(struct node));
   printf("Enter data: ");
   scanf("%d", &newnode->data);
   while(i<pos-1)
   {
      temp = temp->next;
      i++;
   }
   newnode->prev = temp;
   newnode->next = temp->next;
   temp->next = newnode;
   newnode->next->prev = newnode;
}

void insertAfterPosition()
{
   int pos, i=1;
   printf("Enter position: ");
   scanf("%d", &pos);
   struct node *newnode, *temp;
   temp = head;
   newnode = (struct node*) malloc(sizeof(struct node));
   printf("Enter data: ");
   scanf("%d", &newnode->data);
   while(i<pos)
   {
      temp = temp->next;
      i++;
   }
   newnode->prev = temp;
   newnode->next = temp->next;
   temp->next = newnode;
   newnode->next->prev = newnode;
}

void delFromBeg() // POP
{
   struct node *temp;
   if(head == NULL)
   {
      printf("List is empty");
   }
   else
   {
      temp = head;
      head = temp->next;
      head->prev = NULL;
      free(temp);
   }
}

void delFromEnd()
{
   struct node *temp;
   if(head == NULL)
   {
      printf("List is empty");
   }
   else
   {
      temp = tail;
      tail->prev->next = NULL;
      tail = tail->prev;
      free(temp);
   }
}

void delFromPos()
{
   int pos, i = 1;
   struct node *temp;
   temp = head;
   printf("Enter position: ");
   scanf("%d", &pos);
   while(i<pos)
   {
      temp = temp->next;
      i++;
   }
   temp->prev->next = temp->next;
   temp->next->prev = temp->prev;
   free(temp);
}

int main(int argc, char *argv[])
{
   createDLL();
}
