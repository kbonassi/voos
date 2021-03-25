package main

import (
	"fmt"
	"os"
	"strings"
)

// Arquivo é uma variável que armazena o nome do arquivo CSV passado
// como parâmetro para a execução da aplicação
var Arquivo string

// main é a função principal, a porta de entrada para a aplicação
func main() {
	if len(os.Args) != 3 {
		msg := "\n\nUso: voos <arquivo.csv> <interface de consulta>\n\n" +
			"interfaceDeConsulta pode ser: console ou rest\n\n" +
			"Exemplo: voos input-routes.csv console"
		panic(msg)
	}

	Arquivo = os.Args[1]

	// confere a existência do arquivo CSV informado para continuar a execução
	_, err := os.Stat(Arquivo)
	if !os.IsNotExist(err) {
		if strings.ToUpper(os.Args[2]) == "CONSOLE" {
			console()
		} else if strings.ToUpper(os.Args[2]) == "REST" {
			rest()
		} else {
			fmt.Printf("'%s' não é um modo de execução válido. Utilize CONSOLE ou REST.", os.Args[2])
			fmt.Println()
		}
	} else {
		fmt.Printf("O arquivo %s não existe. Não tenho como prosseguir.", Arquivo)
		fmt.Println()
	}
}
