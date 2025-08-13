package controllers

import (
	"cadastro/models"
	"cadastro/utils"
	"fmt"
	"os"
	"strconv"
)

var id = 1 //implementa o id da pessoa iniciando em 1
func CadastarPessoa(pessoas []models.Pessoa) []models.Pessoa {
	var p models.Pessoa
	var opcao string
	tentativas := 3
	for {
		//--Cdastra nome--
		for i := 0; i < tentativas; i++ {
			fmt.Print("Digite o nome: ")
			p.Nome = utils.ReadLine()
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
			idadeStr := utils.ReadLine()
			idade, err := strconv.Atoi(idadeStr)
			if utils.CheckError(err, "Erro ao converter entrada para número:") {
				os.Exit(0)
			}
			if idade > 0 && idade < 150 {
				p.Idade = idade
				break
			} else {
				fmt.Println("Idade inválida! Tente novamente.")
			}
			if i == tentativas-1 {
				fmt.Println("Numero de tentativas atingido. Encerrando o programa...")
				os.Exit(1)
			}

		}
		//--Cadastra CPF--
		for i := 0; i < tentativas; i++ {
			fmt.Print("Digite o CPF: ")
			p.CPF = utils.ReadLine()
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
			alturaStr := utils.ReadLine()
			altura, err := strconv.ParseFloat(alturaStr, 64)
			if utils.CheckError(err, "Erro ao converter entrada para número:") {
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

		p.Id = id
		id++
		//--Adicionando a pessoa no slice
		pessoas = append(pessoas, p)
		fmt.Print("Deseja continuar cadastrando? s/n: ")
		opcao = utils.ReadLine()
		fmt.Println("")
		if opcao == "s" {
			utils.Clear()
			continue
		} else if opcao == "n" {
			utils.Clear()
			break
		} else {
			fmt.Println("Valor digitado inválido.")
			os.Exit(0)
		}

	}
	return pessoas
}
