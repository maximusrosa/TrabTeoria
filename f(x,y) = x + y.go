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

// A:=fst(B) usando C,D (extrai o primeiro componente do par)
operation fst(A,B,C,D){
  1: do clear(A) goto 2
  2: do load(C,B,D) goto 3
  3: if divBy2(C,D) then goto 4 else goto 6
  4: do div2(C,D) goto 5
  5: do inc A goto 3
  6: do clear(C) goto 0
}

// A:=A div 3 usando C (divisão inteira por 3)
operation div3(A,C){
  1: do load(C,A) goto 2
  2: if zero C then goto 0 else goto 3
  3: do dec C goto 4
  4: if zero C then goto 0 else goto 5
  5: do dec C goto 6
  6: if zero C then goto 0 else goto 7
  7: do dec C goto 8
  8: do inc A goto 2
}

// A divisivel_por_3 usando C (teste se múltiplo de 3)
test divBy3(A,C){
  1: do load(C,A) goto 2
  2: if zero C then goto true else goto 3
  3: do dec C goto 4
  4: do inc A goto 5
  5: if zero C then goto false else goto 6
  6: do inc A goto 7
  7: do dec C goto 8
  8: if zero C then goto false else goto 9
  9: do dec C goto 10
 10: do inc A goto 2
}

// A:=snd(B) usando C,D (extrai o segundo componente do par)
operation snd(A,B,C,D){
  1: do clear(A) goto 2
  2: do load(C,B,D) goto 3
  3: if divBy3(C,D) then goto 4 else goto 6
  4: do div3(C,D) goto 5
  5: do inc A goto 3
  6: do clear(C) goto 0
}

// Programa principal
main {
  // extrai o primeiro componente do par
  1: do fst(A,X,C,D) goto 2  
  // extrai o segundo  componente do par
  2: do snd(B,X,C,D) goto 3  
  // coloca em Y a soma dos dois componentes
  3: do soma(Y,A) goto 4      
  4: do soma(Y,B) goto 0      
} 