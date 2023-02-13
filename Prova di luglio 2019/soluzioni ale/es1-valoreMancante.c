// Alessandro Belfiore 975826
#include <stdio.h>
#include <stdlib.h>

int valoreMancante(int *A, int n) {
  if(A[0] != 0) {
    return 0;
  }
  if(A[n-1] != n) {
    return n;
  }
  int sx = 0, dx = n, m;
  while(sx < dx) {
    m = (sx + dx) / 2;
    if((A[m] != m) && (A[m-1] == m-1)) {
      return m;
    } else if(A[m] == m) {
      sx = m + 1;
    } else {
      dx = m;
    }
  }
  return 0;
}

int main( void ) {
  int n,num;
  scanf("%d",&n);
  int *A = malloc( n * sizeof(int) );
  for(int i = 0; i < n; i++) {
    scanf("%d",&num);
    A[i] = num;
  }
  printf("%d\n",valoreMancante(A,n));
  return 0;
}
