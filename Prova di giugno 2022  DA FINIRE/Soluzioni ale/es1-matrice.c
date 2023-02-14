// Alessandro Belfiore 975826
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int calcolaPot(int *a, int m) {
  int num = 0;
  for(int i = 0; i < m; i++) {
    if(a[i] == 1) {
      num = num + ((int)pow(2,i));
    }
  }
  return num;
}

int *inverti(int *a, int m) {
  int *b = calloc( m,sizeof(int) );
  for(int i = 0; i < m; i++) {
    b[i] = a[(m-1) - i];
  }
  return b;
}

int main( void ) {
  int n,m,conta = 0;
  scanf("%d %d",&n,&m);
  int **matr = calloc( n,sizeof(int *) );
  for(int i = 0; i < n; i++) {
    matr[i] = calloc( m,sizeof(int) );
  }
  for(int i = 0; i < n; i++) {
    for(int j = 0; j < m; j++) {
      scanf("%d",&matr[i][j]);
      if(matr[i][j] == 1) {
        conta++;
      }
    }
  }
  printf("%d ",conta);
  int *app = calloc( m,sizeof(int) );
  for(int i = 0; i < n; i++) {
    app = inverti(matr[i],m);
    printf("%d ",calcolaPot(app,m));
  }
  printf("\n");
  return 0;
}
