package main

import "fmt"

func main() {
    var saldo float64
    var opcao int
    var valor float64

    fmt.Println("Digite seu saldo inicial: R$")
    fmt.Scan(&saldo)

    for {
        fmt.Println("Saldo atual: R$", saldo)
        fmt.Println("1. Depositar")
        fmt.Println("2. Sacar")
        fmt.Println("3. Sair")
        fmt.Println("➔ Escolha uma opção:")
        fmt.Scan(&opcao)

        if opcao == 1 {
            fmt.Println("Valor para depositar: R$")
            fmt.Scan(&valor)
            if valor > 0 {
                saldo += valor
                fmt.Println("Depósito de R$", valor, "realizado!")
            } else {
                fmt.Println("Valor inválido!")
            }

        } else if opcao == 2 {
            fmt.Println("Valor para sacar: R$")
            fmt.Scan(&valor)
            if valor > saldo {
                fmt.Println("Saldo insuficiente!")
            } else if valor <= 0 {
                fmt.Println("Valor inválido!")
            } else {
                saldo -= valor
                fmt.Println("Saque de R$", valor, "realizado!")
            }

        } else if opcao == 3 {
            fmt.Println("Encerrando... Saldo final:", saldo)
            break

        } else {
            fmt.Println("Opção inválida! Use 1, 2 ou 3.")
        }
    }
}