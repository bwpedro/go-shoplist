package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func checkFile(err error) bool {
	if err != nil {
		log.Fatal("Arquivo não pode ser criado", err)
		return true
	}
	return false
}

func pegaDados(email, senha *string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("E-mail: ")
	scanner.Scan()
	*email = scanner.Text()

	fmt.Printf("Senha: ")
	scanner.Scan()
	*senha = scanner.Text()
}

func cadastrarLogin(fileLogin *os.File) int {
	var email, senha string
	pegaDados(&email, &senha)

	rand.Seed(time.Now().UTC().UnixNano())
	id := rand.Intn(10)

	fmt.Printf("\nCadastro realizado com sucesso!\n\n")
	fmt.Fprintf(fileLogin, "%d;%s;%s", id, email, senha)

	return id
}

func validaLogin(fileLogin *os.File) {
	var email, senha string
	pegaDados(&email, &senha)
	fmt.Printf("Validando login...\n")
	fmt.Printf("Bem-vindo de volta!\n")
}

func addCarrinho(filePratos, fileLista *os.File) {
	scanner := bufio.NewScanner(os.Stdin)
	var prato string
	mais := true

	fmt.Printf("Os pratos disponíveis são:\n")
	// lista pratos

	fmt.Printf("Adicionar ao carrinho: \n")
	fmt.Printf("Digite 0 para fechar suas compras\n")
	for mais {
		scanner.Scan()
		prato = scanner.Text()
		if prato == "0" {
			mais = false
		} else {
			// lista de pratos para adicionar a lista
			fmt.Fprintln(fileLista, prato)
		}
	}
	fmt.Printf("Pratos adicionados com sucesso!\n")
	return
}

func consultarLista(fileLista *os.File) {
	fmt.Printf("Sua lista possui os pratos: \n")
	//listar pratos da lista
	gerarLista(fileLista)
	return
}

func gerarLista(fileLista *os.File) {
	gerarList := true
	selec := 0
	for gerarList {
		fmt.Printf("Deseja gerar sua lista de compras? (1-Sim/2-Não)\n")
		fmt.Scanf("%d", &selec)
		if selec == 1 {
			gerarList = false
		} else if selec == 2 {
			return
		} else {
			fmt.Println("Opção não disponível")
		}
	}

	// printa lista de ingredientes
	fmt.Printf("Lista gerada com sucesso! Consulte ingredientes.txt\n")
}

func limparCarrinho(fileLista *os.File) {
	fmt.Println("Carrinho limpo!")
}
