// Alessandro Belfiore 975826
#include <stdio.h>

void scambia(int *a, int i, int j) {
  int temp = a[i];
  a[i] = a[j];
  a[j] = temp;
}

void selectionSort(int *a, int n) {
  int k;
  for(int i = 0; i < n - 1; i++) {
    k = i;
    for(int j = i + 1; j < n; j++) {
      if(a[j] < a[k]) {
        k = j;
      }
    }
    scambia(a,i,k);
  }
}

int sceltaGreedy(int *a, int L, int n) {
  selectionSort(a,n);
  int sum = 0,conta = 0;
  for(int i = 0; i < n; i++) {
    sum = sum + a[i];
    if(sum <= L) {
      conta++;
    } else {
      break;
    }
  }
  return conta;
}

int main( void ) {
  int L,n;
  scanf("%d %d",&L,&n);
  int a[n];
  for(int i = 0; i < n; i++) {
    scanf("%d",&a[i]);
  }
  printf("%d\n",sceltaGreedy(a,L,n));
  return 0;
}
