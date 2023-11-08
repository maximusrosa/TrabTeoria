// A:=0 
operation zerar(A){ 
    1: if zero A then goto 0 else goto 2 
    2: do dec A goto 1 
} 

// A:=A+B (soma destrutiva) 
operation soma(A,B){ 
    1: if zero B then goto 0 else goto 2 
    2: do dec B goto 3 
    3: do inc A goto 1 
} 

// A:=A+B usando C (soma conservativa) 
operation soma(A,B,C){
    1: if zero B then goto 5 else goto 2 
    2: do dec B goto 3 
    3: do inc A goto 4 
    4: do inc C goto 1 
    5: if zero C then goto 0 else goto 6
    6: do dec C goto 7 
    7: do inc B goto 5 
} 

// A:=B (atribuicao destrutiva) 
operation atribui(A,B){ 
    1: do zerar(A) goto 2 
    2: do soma(A,B) goto 0 
} 

// A:=B usando C (atribuição conservativa) 
operation atribui(A,B,C){ 
    1: do zerar(A) goto 2 
    2: do soma(A,B,C) goto 0 
}



main {

// A = X + 1
    1: do atribui(A, X, C) goto 2
    2: do inc A goto 3


// X = X x A
//  C:=X;
//  até C=0 faça 
//     (X:= X+A usando D;
//      C:=C-1)

   3: do atribui(C, X) goto 4

   4: do soma(X, A, D) goto 5
   5: do dec C goto 6
   6: if zero C then goto 7 else goto 4

// Y = X  / 2
//  Y:=X;
//  até X=0 faça
//    (X:=C-1;
//     X:=C-1;
//     Y:=Y-1)

   7: do atribui(Y, X, C) goto 8

   8: if zero X then goto 0 else goto 9
   9: do dec X goto 10
  10: do dec X goto 11  
  11: do dec Y goto 8
  
}