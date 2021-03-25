package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// melhorRota é uma estrutura para receber os registros CSV
type melhorRota struct {
	Rota  []string `json:"rota"`
	Custo int      `json:"custo"`
}

// voosHandler é uma função de entrada para responder a API e
// dela rotear para função correta conforme os métodos do
// protocolo HTTP implementados
func voosHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	rota := ""
	if params["rota"] != nil {
		rota = params["rota"][0]
	}

	switch {
	case r.Method == "GET" && len(rota) > 0:
		voosPorRota(w, r, rota)
	case r.Method == "GET":
		voosTodasRotas(w, r)
	case r.Method == "POST":
		voosNovaRota(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Método não implementado no endpoint")
	}
}

// retornaMelhorRota é a função responsável por preparar
// o vetor de resposta de melhor rota, utilizada pelo
// endpoint voosPorRota
func retornaMelhorRota(dict map[string]int, rota []int, custo int, origem string, destino string) melhorRota {
	retArray := []string{fmt.Sprintf("Não existe nenhum caminho de %s para %s", origem, destino)}

	if len(rota) > 1 {
		retArray = nil
		for v := range rota {
			key, ok := mapDict(dict, rota[v])
			if ok {
				retArray = append(retArray, key)
			}
		}
	}

	retorno := melhorRota{Rota: retArray, Custo: custo}
	return retorno
}

// voosPorRota é uma função (chamada pelo endpoint) para o tratamento
// e retorno da escolha da possível melhor rota para um trecho de voo.
// Esta função faz uso do algorítmo de dijkstra para trazer o resultado
func voosPorRota(w http.ResponseWriter, r *http.Request, rota string) {
	dict, matriz, _ := LoadCSV(Arquivo)

	if !validaRota(rota) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Rota inválida. Use XXX-XXX")
	} else {
		origem, destino := splitRoute(rota)
		origem = strings.ToUpper(origem)
		destino = strings.ToUpper(destino)
		if !validaAeroporto(dict, origem) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Aeroporto de origem inválido")
		} else if !validaAeroporto(dict, destino) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Aeroporto de destino inválido")
		} else {
			melhorRota, custo := Dijkstra(matriz, dict[origem], dict[destino])
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json, _ := json.Marshal(retornaMelhorRota(dict, melhorRota, custo, origem, destino))
			fmt.Fprint(w, string(json))
		}
	}
}

// voosTodasRotas é uma função (chamada pelo endpoint) para o
// retorno de todos os trechos de rotas de voo do arquivo CSV.
func voosTodasRotas(w http.ResponseWriter, r *http.Request) {
	_, _, voos := LoadCSV(Arquivo)

	json, _ := json.Marshal(voos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))
}

// voosNovaRota é uma função (chamada pelo endpoint) para a
// inclusão de novas rotas e atualização de custo de rotas
// existentes no arquivo CSV.
func voosNovaRota(w http.ResponseWriter, r *http.Request) {
	var idxV, idxR int
	_, _, voos := LoadCSV(Arquivo)
	var rota []Voo

	err := json.NewDecoder(r.Body).Decode(&rota)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, v := range rota {
		idxR = i
		idxV = null
		for j, w := range voos {
			if v.Origem == w.Origem &&
				v.Destino == w.Destino {
				idxV = j
			}
		}
		if idxV > 0 { // atualizacao de custo, rota pre-existente
			voos[idxV].Custo = rota[idxR].Custo
		} else { // nova rota
			voos = append(voos, rota[idxR])
		}
	}

	retorno := gravaCSV(voos)

	if retorno {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Rotas atualizadas com sucesso")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Problemas ao escrever arquivo .CSV")
	}
}

// rest é uma função para controle da aplicação no modo REST
func rest() {
	http.HandleFunc("/voos/", voosHandler)
	log.Println("Executando a API REST na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
