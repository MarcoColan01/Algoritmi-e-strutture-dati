// Alessandro Belfiore 975826
/* Schema dei processi produttivi rappresentato attraverso un array di lunghezza uguale al numero di processi
lavorativi, dove in posizione array[i], viene inserito un puntatore alla radice dell'albero rappresentante
l'impianto di produzione.*/

/* Il numero di lavoratori è contenuto all'interno del singolo nodo che viene così strutturato:
*/

struct node {
  int num_reparto,num_lavoratori;
  struct node *left;
  struct node *right;
};
typedef struct *Bit_node

int lavoratoriInCollaudo(Bit_node root) {
  if(root ==  NULL) {
    return 0;
  }
  if(root -> left == NULL && root -> right == NULL) {
    return root -> num_lavoratori;
  }
  return lavoratoriInCollaudo(root -> left) + lavoratoriInCollaudo(root -> right);
}

void Algoritmo(int *a, int n) {
  int sum = 0;
  for(int i = 0; i < n; i++) {
    sum = sum + lavoratoriInCollaudo(a[i]);
  }
  printf("I lavoratori dei reparti di collaudo dei diversi impianti di produzione sono :%d\n",sum);
}
