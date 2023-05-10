// Alessandro Belfiore 975826
/*
1)
a) Un puntatore ad una struttura nodo.
b) Viene invocata tante volte quanta è lunga la parola passata come parametro alla funzione build.
c) Se la parola del nodo corrente è composta da un solo carattere.
d) Utilizzando l'aritmetica dei puntatori viene spostato il puntatore ad il carattere successivo della parola.
2)
fstring, data una lista di strutture nodo, aventi per item parole, concatena tali parole, partendo dal fondo
e risalendo la lista, ottenendo come risultato finale la parola presente nell'ultimo
nodo della lista concatenata con i suffissi ogni volta eliminando il primo carattere;
il primo carattere da passare alla funzione deve essere una stringa vuota.
*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct node {
	char *w;
	struct node *next;
};
typedef struct node *Node;

char *fstring( char *w, Node l ){
	if ( l -> next == NULL )
		return strcat( w, l -> w );

	return strcat( fstring( w, l -> next ), l -> w );
}

void fstring_iter( Node l ){
  Node t1 = l,t2;
  while( t1 ) {
    while(t1 -> next != NULL) {
      t2 = t1;
      t1 = t1 -> next;
    }
    printf("%s",t1 -> w);
    if ( t1 -> w[0] != '\0' && t1 -> w[1] == '\0' ) {
      printf("\n");
      return;
    }
    t2 -> next = NULL;
    t1 = l;
  }
}

Node flist( Node l ) {
	if ( l -> w[0] != '\0' && l -> w[1] == '\0' )
		return l;
	Node n = malloc( sizeof(struct node) );
	n -> w = ( l -> w ) + 1;
	n -> next = l;
	return flist( n );
}

Node build( char *w ) {
	Node lista = malloc (sizeof(struct node) );
	lista -> w = w;
	lista -> next = NULL;
	return flist( lista );
}

Node flist2( Node l ) {
	if ( l -> w[0] != '\0' && l -> w[1] == '\0' )
		return l;
	Node n = malloc( sizeof(struct node) );
  n -> w = malloc(20);
	for(int i = 0; i < strlen(l -> w) - 1; i++) {
    n -> w[i] = l -> w[i];
  }
	n -> next = l;
	return flist2( n );
}

Node build2( char *w ) {
	Node lista = malloc (sizeof(struct node) );
	lista -> w = w;
	lista -> next = NULL;
	return flist2( lista );
}

void stampaLista(Node lista) {
	Node t;
	for(t = lista; t != NULL; t = t -> next) {
		printf("%s ",t -> w);
	}
	printf("\n");
}

int main( void ) {
	Node newLista = build("cane");
	stampaLista(newLista);
	char word[20] = "";
  char word1[20] = "";
	printf("%s\n",fstring(word,newLista));
  fstring_iter(newLista);
  printf("%s\n",fstring(word1,build2("hello")));
	return 0;
}
