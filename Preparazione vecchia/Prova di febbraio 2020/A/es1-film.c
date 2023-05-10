// Alessandro Belfiore 975826
#include <stdio.h>
#include <stdlib.h>

int main( void ) {
  int n,conta = 1,max = 1;
  scanf("%d",&n);
  int P[n];
  for(int i = 0; i < n; i++) {
    scanf("%d",&P[i]);
  }
  for(int i = 1; i < n; i++) {
    if(P[i] > P[i-1]) {
      conta++;
      if(conta > max) {
        max = conta;
      }
    } else {
      conta = 1;
    }
  }
  printf("%d\n",max);
  return 0;
}

/*
Costo computazionale = O(n)
*/
