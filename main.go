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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var p Pessoa
	fmt.Println("Sistema de Cadastro de Pessoas")

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

	fmt.Println(p)

}
