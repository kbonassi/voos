package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// splitRoute é uma função que faz o split de rotas origem e destino
// com base em uma string informada
func splitRoute(route string) (string, string) {
	return route[:3], route[4:7]
}

// validateRoute é uma função para validar se a rota informada atende ao
// padrão exigido de 3 letras - 3 letras (XXX-XXX). Utiliza regular expression
// para realizar a validação
func validaRota(rota string) bool {
	match, _ := regexp.MatchString("(?m)^[aA-zZ]{3}(-[aA-zZ]{3})$", rota)

	return match
}

// validaAeroporto é uma função que valida se o aeroporto informado existe
// no arquivo CSV, seja ele de origem ou destino
func validaAeroporto(dict map[string]int, aero string) bool {
	if _, ok := dict[aero]; ok {
		return ok
	}
	return false
}

// mapDict é uma função que retorna a chave de um map baseado em seu valor
// utilizada para sabermos o nome da chave conforme o valor no dicionário de
// aeroportos.
func mapDict(dict map[string]int, valor int) (string, bool) {
	for k, v := range dict {
		if v == valor {
			return k, true
		}
	}
	return "", false
}

// imprimeMelhorRota é uma função que retorna uma string formatada
// da possível rota entre origem e destino informado pelo algoritmo de pesquisa
func imprimeMelhorRota(dict map[string]int, rota []int, custo int, origem string, destino string) string {
	retorno := fmt.Sprintf("Não existe nenhum caminho de %s para %s", origem, destino)

	if len(rota) > 1 {
		retorno = ""
		for v := range rota {
			key, ok := mapDict(dict, rota[v])
			if ok {
				retorno += key
				if v != len(rota)-1 {
					retorno += " - "
				}
			}
		}
		retorno += fmt.Sprintf(" > $%d", custo)
	}
	return retorno
}

// console é uma função para controle da aplicação no modo CLI
func console() {
	dict, matriz, _ := LoadCSV(Arquivo)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("please enter the route: ")
		route, _ := reader.ReadString('\n')
		if !validaRota(route) {
			fmt.Println("Rota inválida. Use XXX-XXX")
		} else {
			origem, destino := splitRoute(route)
			origem = strings.ToUpper(origem)
			destino = strings.ToUpper(destino)
			if !validaAeroporto(dict, origem) {
				fmt.Println("Aeroporto de origem inválido")
			} else if !validaAeroporto(dict, destino) {
				fmt.Println("Aeroporto de destino inválido")
			} else {
				rota, custo := Dijkstra(matriz, dict[origem], dict[destino])
				fmt.Print("best route: ")
				fmt.Println(imprimeMelhorRota(dict, rota, custo, origem, destino))
			}
		}
	}
}
