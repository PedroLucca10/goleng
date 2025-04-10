package main

import (
  "fmt"
)

func main (){
var numeros = [5]int{}
 
fmt.Println("Digite o 1º número: ")
fmt.Scan(&numeros[0])
fmt.Println("Digite o 2º número: ")
fmt.Scan(&numeros[1])
fmt.Println("Digite o 3º número: ")
fmt.Scan(&numeros[2])
fmt.Println("Digite o 4º número: ")
fmt.Scan(&numeros[3])
fmt.Println("Digite o 5º número: ")
fmt.Scan(&numeros[4])
fmt.Println("A soma dos números é: ", numeros[0] + numeros[1] + numeros[2] + numeros[3] + numeros[4])

}