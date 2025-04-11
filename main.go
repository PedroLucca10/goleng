package main

import "fmt"

func sayGreeting(nome string){
    fmt.Println("Olá", nome)
}

func addNumber(numero1 int, numero2 int){
    return numero1 + numero2
}
func main(){
    sayGreeting("Juvelino")
}



func main(){
    sayGreeting("Juvelino")
    resultado := addNumber (10, 20)
    fmt.Println(resultado)
}