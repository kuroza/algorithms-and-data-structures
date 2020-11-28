#include <stdio.h>
#include <stdlib.h>

void main()
{
   struct node *root;
   printf(Pre-order is: );
   preorder(root);
   inOrder(root);
   postOrder(root);
}

void preOrder(struct node *root)
{
   if(root == 0) {
      return;
   }
   printf("%d", root->data);
   preOrder(root->left);
   preOrder(root->right);
}

void inOrder(struct node *root)
{
   if(root == 0)
   {
      return;
   }
   inOrder(root->left);
   printf("%d", root->data);
   inOrder(root->right);
}

void postOrder(struct node *root)
{
   if(root == 0)
   {
      return;
   }
   postOrder(root->left);
   postOrder(root->right);
   printf("%d", root->data);
}
