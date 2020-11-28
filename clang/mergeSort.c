// Merge sort time complexity O(n log n)

mergeSort(A, lb, ub)
{
   if(lb < ub)
   {
      mid = (lb + ub) / 2;
      mergeSort(A, lb, mid);
      mergeSort(A, mid+1, ub);
      mergeSort(A, lb, mid, ub);
   }
}

merge(A, lb, mid, ub)
{
   i = lb;
   j = mid+1;
   k = lb;

   while(i<=mid && j<=ub)
   {
      if(a[i] <= a[j])
      {
         b[k] = a[i];
         i++;
      }
      else
      {
         b[k] = a[j];
         j++;
      }
      k++;
   }

   if(i>mid)
   {
      while(j <= ub)
      {
         b[k] = a[j];
         j++;
         k++;
      }
   }
   else
   {
      while)i <= mid)
      {
         b[k] = a[i];
         i++;
         k++;
      }
   }

   for(k=lb; k<=ub; k++)
   {
      a[k] = b[k];
   }
}
