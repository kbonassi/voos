package main

import (
	"testing"
)

func TestConsole(t *testing.T) {
	dict, matriz, _ := LoadCSV("input-file.csv")
	origem := "GRU"
	destino := "CDG"
	rota, custo := Dijkstra(matriz, dict[origem], dict[destino])
	resObtido := imprimeMelhorRota(dict, rota, custo, origem, destino)

	// confere se o retorno é o esperado conforme arquivo CSV para testes.
	resEsperado := `GRU - BRC - SCL - ORL - CDG > $40`
	if resObtido != resEsperado {
		t.Errorf("Resposta não esperada:\n retornou %v\nesperava %v",
			resObtido, resEsperado)
	}
}
