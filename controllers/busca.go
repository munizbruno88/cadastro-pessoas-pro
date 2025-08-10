package controllers

import (
	"cadastro/models"
	"cadastro/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BuscarPessoas(pessoas *[]models.Pessoa) {
	var opcao int
	for {
		fmt.Println("Selecione as opções abaixo.")
		fmt.Println("1 - Busca por nome:")
		fmt.Println("2 - Busca por ID:")
		fmt.Println("0 - Voltar")
		fmt.Print("Digite uma opção: ")
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			utils.Clear()
			fmt.Println("Opção inválida! Digite apenas numeros.")
			fmt.Println("")
			continue
		}
		fmt.Println("")
		switch opcao {
		case 1:
			utils.Clear()
			fmt.Print("Digite o nome: ")
			nome := strings.ToLower(utils.ReadLine())
			encontrado := false
			for _, p := range *pessoas {
				if strings.Contains(strings.ToLower(p.Nome), nome) {
					utils.Clear()
					encontrado = true
					BuscaMenu(p, pessoas)
					return

				}
			}
			if !encontrado {
				fmt.Println("Pessoa não encontrada.")
				utils.Clear()
			}
		case 2:
			utils.Clear()
			fmt.Print("Digite o ID: ")
			iDStr := utils.ReadLine()
			iD, err := strconv.Atoi(iDStr)
			if err != nil {
				os.Exit(0)
			}
			encontrado := false
			for _, p := range *pessoas {
				if p.Id == iD {
					utils.Clear()
					encontrado = true
					BuscaMenu(p, pessoas)
					return

				}
			}
			if !encontrado {
				fmt.Println("Pessoa não encontrada.")
				utils.Clear()
			}
		case 0:
			utils.Clear()
			return
		default:
			utils.Clear()
			fmt.Println("Opção inválida! Tente novamente.")
			fmt.Println("")
			continue

		}
	}

}

func Editar(pessoas *[]models.Pessoa, id int) {
	utils.Clear()
	fmt.Print("Digite o novo nome: ")
	novoNome := utils.ReadLine()
	utils.Clear()
	for i := range *pessoas {
		if (*pessoas)[i].Id == id {
			(*pessoas)[i].Nome = novoNome
			fmt.Printf("ID: %d\nNome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm\n\n", (*pessoas)[i].Id, (*pessoas)[i].Nome, (*pessoas)[i].Idade, (*pessoas)[i].CPF, (*pessoas)[i].Altura)
		}
	}

}

func deletarPessoa(pessoas *[]models.Pessoa, id int) {
	for i := range *pessoas {
		if (*pessoas)[i].Id == id {
			*pessoas = append((*pessoas)[:i], (*pessoas)[i+1:]...)
			fmt.Println("Pessoa removida com sucesso!")
			fmt.Println("")
			return
		}
	}
	fmt.Println("Pessoa não encontrada!")
}

func BuscaMenu(p models.Pessoa, pessoas *[]models.Pessoa) {
	fmt.Printf("ID: %d\nNome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm\n\n", p.Id, p.Nome, p.Idade, p.CPF, p.Altura)
	fmt.Println("1 - Editar")
	fmt.Println("2 - Deletar")
	fmt.Println("0 - Voltar")
	fmt.Print("Digite uma opção: ")
	opcaoBuscaStr := utils.ReadLine()
	opcaoBusca, err := strconv.Atoi(opcaoBuscaStr)
	if err != nil {
		os.Exit(0)
	}
	switch opcaoBusca {
	case 1:
		Editar(pessoas, p.Id)
	case 2:
		fmt.Print("Tem certeza que deseja excluir? s/n: ")
		sN := utils.ReadLine()
		switch sN {
		case "s":
			deletarPessoa(pessoas, p.Id)
		case "n":
			return
		}
	}
}
