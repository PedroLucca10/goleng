package main

import "fmt"

func main() {
    estoque := map[string]int{
        "Coxinha":        10,
        "Pão de queijo":  15,
        "Refrigerante":   20,
    }
    estoque["Coxinha"] -= 2    
    estoque["Pão de queijo"] -= 1 

    fmt.Println("Estoque atualizado:")
    for produto, quantidade := range estoque {
       fmt.Println(produto + ":", quantidade)
    }
}
