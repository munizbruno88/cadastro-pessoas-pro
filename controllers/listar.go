package controllers

import (
	"cadastro/models"
	"cadastro/utils"
	"fmt"
	"os"
	"strconv"
)

func ListarPessoas(pessoas []models.Pessoa) {
	if len(pessoas) == 0 {
		fmt.Println("Não há pessoas cadastradas!")
	} else {
		fmt.Println("\n##--PESSOAS CADASTRADAS--##")
		for _, p := range pessoas {
			fmt.Printf("ID: %d\nNome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm\n\n", p.Id, p.Nome, p.Idade, p.CPF, p.Altura)

		}
		fmt.Print("Digite 0 para voltar: ")
		sairStr := utils.ReadLine()
		sair, err := strconv.Atoi(sairStr)
		if err != nil {
			os.Exit(0)
		}
		if sair == 0 {
			utils.Clear()
			return
		} else {
			utils.Clear()
			fmt.Println("Opção inválida! Encerrando programa...")
			os.Exit(0)
		}
	}
}
