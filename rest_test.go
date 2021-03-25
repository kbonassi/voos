package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestGetVoosTodasRotas(t *testing.T) {
	res, err := http.Get("http://localhost:3000/voos/")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	corpo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	resObtido := fmt.Sprintf("%s", corpo)

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Status não esperado: retornou %v esperava %v",
			status, http.StatusOK)
	}

	// confere se o retorno é o esperado conforme arquivo CSV para testes.
	resEsperado := `[{"origem":"GRU","destino":"BRC","custo":10},{"origem":"BRC","destino":"SCL","custo":5},{"origem":"GRU","destino":"CDG","custo":75},{"origem":"GRU","destino":"SCL","custo":20},{"origem":"GRU","destino":"ORL","custo":56},{"origem":"ORL","destino":"CDG","custo":5},{"origem":"SCL","destino":"ORL","custo":20}]`
	if resObtido != resEsperado {
		t.Errorf("Resposta não esperada:\n retornou %v\nesperava %v",
			resObtido, resEsperado)
	}
}

func TestGetVoosPorRota(t *testing.T) {
	res, err := http.Get("http://localhost:3000/voos/?rota=GRU-CDG")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	corpo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	resObtido := fmt.Sprintf("%s", corpo)

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Status não esperado: retornou %v esperava %v",
			status, http.StatusOK)
	}

	// confere se o retorno é o esperado conforme arquivo CSV para testes.
	resEsperado := `{"rota":["GRU","BRC","SCL","ORL","CDG"],"custo":40}`
	if resObtido != resEsperado {
		t.Errorf("Resposta não esperada:\n retornou %v\nesperava %v",
			resObtido, resEsperado)
	}
}

func TestPostVoosNovaRota(t *testing.T) {
	url := "http://localhost:3000/voos/"
	var json = []byte(`[{"origem": "ORL","destino": "CDG","custo": 5}]`)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	corpo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	resObtido := fmt.Sprintf("%s", corpo)

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Status não esperado: retornou %v esperava %v",
			status, http.StatusOK)
	}

	// confere se o retorno é o esperado conforme arquivo CSV para testes.
	resEsperado := `Rotas atualizadas com sucesso`
	if resObtido != resEsperado {
		t.Errorf("Resposta não esperada:\n retornou %v\nesperava %v",
			resObtido, resEsperado)
	}
}
