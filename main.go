package main

import "fmt"

func main (){
   var idade = [4]int{15,15,16,16}
     nomes := [4]string {"Bruno", "Yan", "Vini", "Eduardo"}
     fmt.Println(idade)
     fmt.Println(nomes)
     nomes[2] = "Pedro"
     fmt.Println(nomes)
}
