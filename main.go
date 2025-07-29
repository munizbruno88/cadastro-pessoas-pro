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
	var pessoas []Pessoa
	fmt.Println("Sistema de Cadastro de Pessoas")
	for {
		//--Cdastra nome--
		fmt.Print("Digite o nome: ")
		scanner.Scan()
		p.Nome = scanner.Text()

		//--Cadastra idade--
		fmt.Print("Digite a idade: ")
		scanner.Scan()
		idadeStr := scanner.Text()
		idade, err := strconv.Atoi(idadeStr)
		if err != nil {
			os.Exit(0)
		}
		p.Idade = idade

		//--Cadastra CPF--
		fmt.Print("Digite o CPF: ")
		scanner.Scan()
		p.CPF = scanner.Text()

		//--Cadastra Altura--
		fmt.Print("Digite a altura: ")
		scanner.Scan()
		alturaStr := scanner.Text()
		altura, err := strconv.ParseFloat(alturaStr, 64)
		if err != nil {
			os.Exit(0)
		}
		p.Altura = altura
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
