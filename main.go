package main
import (
    "fmt"
  )

 func main () {
     alunoIdade := make(map[string]int)
     alunoIdade["Bruno"] = 15
     alunoIdade["Pedro"] = 16
     alunoIdade["Fabiano"] = 40
     alunoIdade["Cássio"] = 37
     fmt.Println("Idade do Bruno",alunoIdade["Bruno"])

       notasAlunos := make(map[string]float) {
        "Bruno" : 9.7,
        "Pedro" : 10,
        "Fabiano" : 8.5,
        "Cássio" : 6.0,
       }
       for nome.nota := range notasAlunos{
        fmt.Printf("%s tirou notas %.1f \n", k, v)
       }

   }