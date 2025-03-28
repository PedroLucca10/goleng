package main

import "fmt"

func main (){
var usuario string
var senha string
fmt.Print("Nome de usuário: ", usuario)
fmt.Scan(&usuario)
fmt.Print("Senha: ", senha)
fmt.Scan(&senha)
if  usuario == "admin" && senha == "1234" {
     fmt.Println("Acesso permitido, Bem vindo lindão") 
     } else {
       fmt.Println("Acesso negado") 
}
	 }
	