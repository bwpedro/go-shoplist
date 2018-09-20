package main

import (
	"fmt"
	"os"
)

func main() {
	var selec, cadastro int
	var sair, erro bool
	sair = false

	fileLogin, err := os.Create("login.txt")
	filePratos, err := os.Create("pratos.txt")
	fileIngredientes, err := os.Create("ingredientes.txt")
	fileLista, err := os.Create("lista.txt")

	erro = checkFile(err)
	if erro {
		return
	}

	for !sair {
		fmt.Printf("\nBem vindo ao SHOPLIST\nVocê já possui cadastro? (1-Sim/2-Não)")
		fmt.Scanf("%d", &cadastro)
		if cadastro == 1 {
			validaLogin(fileLogin)
			sair = true
		} else if cadastro == 2 {
			cadastrarLogin(fileLogin)
			sair = true
		} else {
			fmt.Println("Opção não disponvível")
		}
	}
	sair = false

	for !sair {
		fmt.Printf("\n\n##### SHOPLIST #####\n\n")
		fmt.Printf("Selecione uma Opção:\n\n1) Adicionar ao carrinho\n2) Consultar lista\n3) Gerar lista\n4) Limpar carrinho\n5) Sair\n")
		fmt.Scanf("%d", &selec)

		switch selec {
		case 1:
			fmt.Printf("\n## ADICIONAR PRATOS AO CARRINHO ##\n\n")
			addCarrinho(filePratos, fileLista)
			break
		case 2:
			fmt.Printf("\n## CONSULTAR LISTA ##\n\n")
			consultarLista(fileLista)
			break
		case 3:
			fmt.Printf("\n## GERAR LISTA ##\n\n")
			gerarLista(fileLista)
			break
		case 4:
			fmt.Printf("\n## LIMPAR CARRINHO ##\n\n")
			limparCarrinho(fileLista)
			break
		case 5:
			sair = true
			fmt.Println("Saindo do programa...")
			break
		default:
			fmt.Println("Opção não disponível")
			break
		}
	}

	defer fileLogin.Close()
	defer filePratos.Close()
	defer fileIngredientes.Close()
	defer fileLista.Close()
}
