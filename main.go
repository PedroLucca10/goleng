package main

import "fmt"

func main (){
   var idade = [4]int{15,15,16,16}
     nomes := [4]string {"Bruno", "Yan", "Vini", "Eduardo"}
     fmt.Println(idade)
     fmt.Println(nomes)
     nomes[2] = "Pedro"
     fmt.Println(nomes)

     var score = []int{100, 200, 300, 400}
     fmt.Println(score)
     score[1] = 2
     fmt.Println(score, len(score), cap(score))
     rangeOne := score[1:3]
     fmt.Println(rangeOne)
     rangeTwo := score[2:]
     fmt.Println(rangeTwo)
     rangeThree := score [:3]
     fmt.Println(rangeThree)


     var superherois = []string{"Batman", "Homem-aranha", "Ben 10"}
     fmt.Println(superherois)
     superherois = append(superherois, "Capitão América")
     fmt.Println(superherois, len(superherois), cap(superherois))
}
