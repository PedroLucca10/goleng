package main

import "fmt"

var saldo float64

func main() {
    fmt.Println("Digite seu saldo inicial: R$")
    fmt.Scan(&saldo)

    for {
        exibirMenu()
        opcao := lerOpcao()

        if opcao == 1 {
            depositar()
        } else if opcao == 2 {
            sacar()
        } else if opcao == 3 {
            encerrar()
            break
        } else {
            fmt.Println("Opção inválida! Use 1, 2 ou 3.")
        }
    }
}

func exibirMenu() {
    fmt.Println("Saldo atual: R$", saldo)
    fmt.Println("1. Depositar")
    fmt.Println("2. Sacar")
    fmt.Println("3. Sair")
    fmt.Println("Escolha uma opção:")
}

func lerOpcao() int {
    var opcao int
    fmt.Scan(&opcao)
    return opcao
}

func depositar() {
    fmt.Println("Valor para depositar: R$")
    var valor float64
    fmt.Scan(&valor)

    if valor > 0 {
        saldo += valor
        fmt.Println("Depósito de R$", valor, "realizado!")
    } else {
        fmt.Println("Valor inválido!")
    }
}

func sacar() {
    fmt.Println("Valor para sacar: R$")
    var valor float64
    fmt.Scan(&valor)

    if valor > saldo {
        fmt.Println("Saldo insuficiente!")
    } else if valor <= 0 {
        fmt.Println("Valor inválido!")
    } else {
        saldo -= valor
        fmt.Println("Saque de R$", valor, "realizado!")
    }
}

func encerrar() {
    fmt.Println("Encerrando... Saldo final:", saldo)
}