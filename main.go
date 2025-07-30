package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pessoa struct {
	Nome   string
	Idade  int
	CPF    string
	Altura float64
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	CadastarPessoa(Pessoa{})

}

func CadastarPessoa(p Pessoa) {
	var opcao string
	tentativas := 3
	var pessoas []Pessoa
	fmt.Println("Sistema de Cadastro de Pessoas")
	for {
		//--Cdastra nome--
		for i := 0; i < tentativas; i++ {
			fmt.Print("Digite o nome: ")
			scanner.Scan()
			p.Nome = scanner.Text()
			if len(p.Nome) > 2 {
				break
			} else {
				fmt.Println("Nome inválido! Tente novamente.")
			}
			if i == tentativas-1 {
				fmt.Println("Numero de tentativas atingido. Encerrando o programa...")
				os.Exit(1)
			}
		}
		//--Cadastra idade--
		for i := 0; i < tentativas; i++ {
			fmt.Print("Digite a idade: ")
			scanner.Scan()
			idadeStr := scanner.Text()
			idade, err := strconv.Atoi(idadeStr)
			if err != nil {
				os.Exit(0)
			}
			if idade > 0 && idade < 150 {
				break
			} else {
				fmt.Println("Idade inválida! Tente novamente.")
			}
			if i == tentativas-1 {
				fmt.Println("Numero de tentativas atingido. Encerrando o programa...")
				os.Exit(1)
			}
			p.Idade = idade
		}
		//--Cadastra CPF--
		for i := 0; i < tentativas; i++ {
			fmt.Print("Digite o CPF: ")
			scanner.Scan()
			p.CPF = scanner.Text()
			if len(p.CPF) == 11 {
				break
			} else {
				fmt.Println("CPF inválido! O CPF possui 11 digitos. Tente novamente.")
			}
			if i == tentativas-1 {
				fmt.Println("Numero de tentativas atingido. Encerrando o programa...")
				os.Exit(1)
			}
		}
		//--Cadastra Altura--
		for i := 0; i < tentativas; i++ {
			fmt.Print("Digite a altura: ")
			scanner.Scan()
			alturaStr := scanner.Text()
			altura, err := strconv.ParseFloat(alturaStr, 64)
			if err != nil {
				os.Exit(0)
			}
			p.Altura = altura
			if p.Altura > 0.5 && p.Altura < 2.5 {
				break
			} else {
				fmt.Println("Altura inválida! Tente novamente.")
			}
			if i == tentativas-1 {
				fmt.Println("Numero de tentativas atingido. Encerrando o programa...")
				os.Exit(1)
			}
		}

		//--Adicionando a pessoa no slice
		pessoas = append(pessoas, p)
		fmt.Print("Deseja continuar cadastrando? s/n: ")
		scanner.Scan()
		opcao = scanner.Text()

		if opcao == "s" {
			continue
		} else if opcao == "n" {
			//--lendo os dados do slice
			for _, p := range pessoas {
				fmt.Println("\nPessoas Cadastradas")
				fmt.Printf("Nome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm ", p.Nome, p.Idade, p.CPF, p.Altura)
				os.Exit(0)
			}
		}

	}
}
