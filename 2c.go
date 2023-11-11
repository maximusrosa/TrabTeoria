// limpa variável
operation clear(A){
 1: if zero A then goto 0 else goto 2
 2: do dec A goto 1
}

// A:=A+B (soma destrutiva)
operation soma(A,B){
  1: if zero B then goto 0 else goto 2
  2: do dec B goto 3
  3: do inc A goto 1
}

// A:=A+B (soma não-destrutiva)
operation soma(A,B,C){
  1: if zero B then goto 5 else goto 2
  2: do dec B goto 3
  3: do inc A goto 4
  4: do inc C goto 1
  5: if zero C then goto 0 else goto 6
  6: do dec C goto 7
  7: do inc B goto 5
}

// A:=B (atribuição destrutiva)
operation load(A,B){
  1: do clear(A) goto 2
  2: do soma(A,B) goto 0
}

// A:=B usando C (atribuição não-destrutiva)
operation load(A,B,C){
  1: do clear(A) goto 2
  2: do soma(A,B,C) goto 0
}

// A:=A div 2 usando C (divisão inteira por 2)
operation div2(A,C){
  1: do load(C,A) goto 2
  2: if zero C then goto 0 else goto 3
  3: do dec C goto 4
  4: if zero C then goto 0 else goto 5
  5: do dec C goto 6
  6: do inc A goto 2
}

// A divisivel_por_2 usando C (teste se multiplo de 2)
test divBy2(A,C){
  1: do load(C,A) goto 2
  2: if zero C then goto true else goto 3
  3: do dec C goto 4
  4: do inc A goto 5
  5: if zero C then goto false else goto 6
  6: do inc A goto 7
  7: do dec C goto 2
}


// codificação: P == 2^a (2b - 1)

// A:=fst(P) usando B,C (extrai o primeiro componente do par)
operation fst(A,P,B,C){

  1: do clear(A) goto 2

  // pra não perder o valor de P:
  2: do load(B,P,C) goto 3

  3: if divBy2(B,C) then goto 4 else goto 0
  4: do div2(B,C) goto 5
  5: do inc A goto 3

  // após essa operação:
  // A := a
  // B := 2b -1
  // P == 2^a (2b - 1)
}

// B:=snd(P) usando A,C (extrai o segundo componente do par)
operation snd(B,P,A,C){

  1: do fst(A,P,B,C) goto 2

  2: do dec B goto 3
  3: do div2(B, C) goto 0

  //4: do clear(A) goto 0

  // após essa operação:
  // A := 0
  // B := b
  // P == 2^a (2b - 1) 
}

operation decod(P, A, B, C){
  // extrai o primeiro componente do par
  1: do fst(A,P,B,C) goto 2  
  // extrai o segundo  componente do par
  2: do snd(B,P,A,C) goto 0
}

// A:=AxB usando C,D (multiplicação destrutiva)
operation mult(A, B, C, D){

// A = A x B
//  C:=A;
//  até C=0 faça 
//     (A:= A+B usando D;
//      C:=C-1)

  1: do load(C, A) goto 2

  2: do soma(A, B, D) goto 3
  3: do dec C goto 4
  4: if zero C then goto 0 else goto 2

}

operation foo(P, A, B, C){
  1: do decod(P, A, B, C) goto 2

  2: do inc A goto 3
  3: do div2(B,C) goto 4

  // Saída no registrador A
  4: do mult(A, B, C, D) goto 0
  
}

// Programa principal
main {
  1: do decod(X, A, B, C) goto 2

  2: do inc A goto 3
  3: do div2(B,C) goto 4

  //1: do foo(X, A, B, C) goto 2
  //2: load(Y, A) goto 0


} 
