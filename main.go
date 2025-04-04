package main

import (
  "fmt"
  "strings"
  "sort"
)

func main (){
  greeting := "Hello my guys"
  fmt.Println(strings.Contains(greeting, "friends"))
  fmt.Println(strings.ReplaceAll(greeting, "my", "mine"))
  fmt.Println(strings.ToUpper(greeting))
  fmt.Println(strings.Index(greeting, "my"))
  fmt.Println(strings.Split(greeting," "))
  ages := []int {50, 80, 10}
  sort.Ints(ages)
  fmt.Println(ages)
  index := sort.SearchInts(ages, 50)
  fmt.Println(index)
  names := []string{"Pedro", "Eduardo", "Bruno"}
  sort.Strings(names)
  fmt.Println(names)
  fmt.Println(sort.SearchStrings(names, "Pedro"))
/*
  x := 0
  for x <= 5 {
    fmt.Println(x)
    x++
}
for 1 :=0; i <5; i++{
  fmt.Println("for 2: ", i)
}

for i:= 0; i <len(names); i++{
  fmt.Println(names[i])
}

fot index, value := range names{
  fmt.Println("O indicio é", index, "e o valor", value)
} */

for index, value := range ages {
  fmt.Println("O indice é", index, "O valor é", value)
  }
  var
superherois := []string("Homen-Aranha", "Hulk", "Homem de ferro")

for index, value range superherois {
  fmt.Println("O Número do herói:", index, "O nome do heroi", value)
}
for i:=0; i <len(superherois); i++{
  fmt.Printl("O numero do heroi", i, "e o heroi", superherois[i])
}
    

}
