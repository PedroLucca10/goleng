package main

import(
  "fmt"
)
var nota1 float64
var nota2 float64

func analisarNotas()(float64, string) {
  fmt.Println("Qual sua 1ª nota: ")
  fmt.Scan(&nota1)
  fmt.Println("Qual sua 2ª nota: ")
  fmt.Scan(&nota2)
  var media = (nota1 + nota2) / 2
  if media < 6 {
    return media, "Reprovado"
  } else if media < 0 || media > 10{
    return media, "Nota inválida"
  } else {
    return media, "Aprovado"
  }
}

  func main() {
    media, resultado := analisarNotas()
    fmt.Println("Média: ", media)
    fmt.Println("Resultado: ", resultado)

  }
  
  