package main

import (
	"cadastro/controllers"
	"cadastro/models"
	"cadastro/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var pessoas []models.Pessoa
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("###--SISTEMA DE CADASTRO DE PESSOAS--##")
		fmt.Println("1 - Cadastrar Pessoas")
		fmt.Println("2 - Listar Pessoas")
		fmt.Println("3 - Buscar Pessoas")
		fmt.Println("0 - Sair")
		fmt.Println("-----------------------------------------")
		fmt.Print("Digite um comando: ")
		opcaostr := utils.ReadLine()
		opcao, err := strconv.Atoi(opcaostr)
		if utils.CheckError(err, "Erro ao converter entrada para número:") {
			os.Exit(0)
		}
		fmt.Println("")
		switch opcao {
		case 1:
			utils.Clear()
			pessoas = controllers.CadastarPessoa(pessoas)

		case 2:
			utils.Clear()
			controllers.ListarPessoas(pessoas)

		case 3:
			utils.Clear()
			controllers.BuscarPessoas(&pessoas)

		case 0:
			os.Exit(0)
		default:
			utils.Clear()
			fmt.Println("Numero inválido! Tente novamente.")
			continue

		}
	}
}
