package main

import (
  "fmt"
)

func main (){
age := 45
fmt.Println(age <= 50)
fmt.Println(age >= 50)
fmt.Println(age == 50)
fmt.Println(age != 50)

if age < 50 {
  fmt.Println("menor que 30 anos")
} else if age < 40 {
  fmt.Println("Menor que 40 anos")
} else {
  fmt.Println("Não é menor que 40 anos")
}

names := []string{"Hugo Souza", "Carrillo", "Garro", "Memphis", "Alberto","Ramon Diaz"}

for index, value := range names {
  if index== 1 {
    fmt.Println("Continue após a posição", index, "E valor", value)
    continue
  }
  if index > 2 {
    fmt.Println("Sair após", index)
  }
  fmt.Println("valor: ", value)
}
}

