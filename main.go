package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pessoa struct {
	Id     int
	Nome   string
	Idade  int
	CPF    string
	Altura float64
}

var id = 1 //implementa o id da pessoa iniciando em 1
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var opcao int
	var pessoas []Pessoa
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("###--SISTEMA DE CADASTRO DE PESSOAS--##")
		fmt.Println("1 - Cadastrar Pessoas")
		fmt.Println("2 - Listar Pessoas")
		fmt.Println("3 - Buscar Pessoas")
		fmt.Println("0 - Sair")
		fmt.Println("-----------------------------------------")
		fmt.Print("Digite um comando: ")
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Opção inválida! Digite apenas numeros.")
			break
		}
		fmt.Println("")
		switch opcao {
		case 1:
			clear()
			pessoas = CadastarPessoa(pessoas)

		case 2:
			clear()
			ListarPessoas(pessoas)

		case 3:
			clear()
			BuscarPessoas(pessoas)

		case 0:
			os.Exit(0)
		default:
			clear()
			fmt.Println("Numero inválido! Tente novamente.")
			continue

		}
	}
}

func CadastarPessoa(pessoas []Pessoa) []Pessoa {
	var p Pessoa
	var opcao string
	tentativas := 3
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

		p.Id = id
		id++
		//--Adicionando a pessoa no slice
		pessoas = append(pessoas, p)
		clear()
		fmt.Print("Deseja continuar cadastrando? s/n: ")
		scanner.Scan()
		opcao = scanner.Text()
		fmt.Println("")
		if opcao == "s" {
			clear()
			continue
		} else if opcao == "n" {
			//fmt.Println(pessoas)
			clear()
			break
		} else {
			fmt.Println("Valor digitado inválido.")
			os.Exit(0)
		}

	}
	return pessoas
}

func ListarPessoas(pessoas []Pessoa) {
	if len(pessoas) == 0 {
		fmt.Println("Não há pessoas cadastradas!")
	} else {
		fmt.Println("\n##--PESSOAS CADASTRADAS--##")
		for _, p := range pessoas {
			fmt.Printf("ID: %d\nNome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm\n\n", p.Id, p.Nome, p.Idade, p.CPF, p.Altura)

		}
		fmt.Print("Digite 0 para voltar: ")
		scanner.Scan()
		sairStr := scanner.Text()
		sair, err := strconv.Atoi(sairStr)
		if err != nil {
			os.Exit(0)
		}
		if sair == 0 {
			clear()
			return
		} else {
			clear()
			fmt.Println("Opção inválida! Encerrando programa...")
			os.Exit(0)
		}
	}
}

func BuscarPessoas(pessoas []Pessoa) {
	var opcao int
	for {
		fmt.Println("Selecione as opções abaixo.")
		fmt.Println("1 - Busca por nome:")
		fmt.Println("2 - Busca por ID:")
		fmt.Println("0 - Voltar")
		fmt.Print("Digite uma opção: ")
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			clear()
			fmt.Println("Opção inválida! Digite apenas numeros.")
			fmt.Println("")
			continue
		}
		fmt.Println("")
		switch opcao {
		case 1:
			clear()
			fmt.Print("Digite o nome: ")
			scanner.Scan()
			nome := strings.ToLower(scanner.Text())

			encontrado := false
			for _, p := range pessoas {
				if strings.Contains(strings.ToLower(p.Nome), nome) {
					fmt.Printf("ID: %d\nNome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm\n\n", p.Id, p.Nome, p.Idade, p.CPF, p.Altura)
					Editar(&pessoas, p.Id)
					encontrado = true
				}
			}
			if !encontrado {
				fmt.Println("Pessoa não encontrada.")

			}
		case 2:
			fmt.Print("Digite o ID: ")
			scanner.Scan()
			iDStr := scanner.Text()
			iD, err := strconv.Atoi(iDStr)
			if err != nil {
				os.Exit(0)
			}
			encontrado := false
			for _, p := range pessoas {
				if p.Id == iD {
					fmt.Printf("ID: %d\nNome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm\n\n", p.Id, p.Nome, p.Idade, p.CPF, p.Altura)
					Editar(&pessoas, p.Id)
					encontrado = true
				}
			}
			if !encontrado {
				fmt.Println("Pessoa não encontrada.")
			}
		case 0:
			return
		default:
			clear()
			fmt.Println("Opção inválida! Tente novamente.")
			fmt.Println("")
			continue

		}
	}

}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func Editar(pessoas *[]Pessoa, id int) {
	fmt.Print("Deseja Editar o nome? s/n: ")
	scanner.Scan()
	sN := scanner.Text()
	if sN == "s" {
		fmt.Print("Digite o novo nome: ")
		scanner.Scan()
		novoNome := scanner.Text()
		for i := range *pessoas {
			if (*pessoas)[i].Id == id {
				(*pessoas)[i].Nome = novoNome
				fmt.Printf("ID: %d\nNome: %s\nIdade: %danos\nCPF: %s\nAltura: %.2fm\n\n", (*pessoas)[i].Id, (*pessoas)[i].Nome, (*pessoas)[i].Idade, (*pessoas)[i].CPF, (*pessoas)[i].Altura)
			}
		}

	}

}
