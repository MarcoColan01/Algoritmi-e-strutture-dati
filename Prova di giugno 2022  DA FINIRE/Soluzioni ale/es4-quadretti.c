// Alessandro Belfiore 975826
#include <stdio.h>
#include <stdlib.h>

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
    }
  }
  int **dinamica = calloc( n,sizeof(int *) );
  for(int i = 0; i < n; i++) {
    dinamica[i] = calloc( m,sizeof(int) );
  }
  for(int i = n - 1; i >= 0; i--) {
    for(int j = m - 1; j >= 0; j--) {
      if(i == n - 1 && matr[i][j] != 1) {
        dinamica[i][j] = 1;
      }
      if(j == m - 1 && matr[i][j] != 1) {
        dinamica[i][j] = 1;
      }
      if(i != n - 1 && j != m - 1 && matr[i][j] != 1) {
        dinamica[i][j] = dinamica[i+1][j] + dinamica[i][j+1];
      }
    }
  }
  printf("\n");
  for(int i = 0; i < n; i++) {
    for(int j = 0; j < m; j++) {
      printf("%d ",dinamica[i][j]);
    }
    printf("\n");
  }
  printf("%d\n",dinamica[0][0]);
  return 0;
}
