package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Voo é uma estrutura para receber os registros CSV
type Voo struct {
	Origem  string `json:"origem"`
	Destino string `json:"destino"`
	Custo   int    `json:"custo"`
}

// checkErr é uma função para tratamento de erro
func checkErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

// LoadCSV é uma função para carregar o arquivo CSV e gerar a
// matriz de adjacencia do grafo
func LoadCSV(fileName string) (map[string]int, [][]int, []Voo) {
	var count = 0
	var origem = ""
	var destino = ""
	var custo = 0

	voos := make([]Voo, 0, 100)
	dict := make(map[string]int)

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("ERRO: ", err.Error())
	}

	defer file.Close()

	// prepara para a leitura do arquivo CSV
	reader := csv.NewReader(bufio.NewReader(file))
	// define o delimitador dos registros CSV como sendo ','
	reader.Comma = ','

	for {
		// realiza a leitura de cada linha do arquivo CSV
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			checkErr(err)
		}

		// carrega o slice de voos
		origem = record[0]
		destino = record[1]
		custo, _ = strconv.Atoi(record[2])
		voos = append(voos, Voo{
			Origem:  origem,
			Destino: destino,
			Custo:   custo,
		})

		// monta o dicionário de aeroportos, verificando origem e destino
		if _, found := dict[origem]; !found {
			dict[origem] = count
			count++
		}
		if _, found := dict[destino]; !found {
			dict[destino] = count
			count++
		}
	}

	// inicializa a matriz (slice) com os voos
	var tamanhoDicionario = len(dict)

	matriz := make([][]int, 0, 100)
	for i := 0; i < tamanhoDicionario; i++ {
		linha := make([]int, tamanhoDicionario)
		for j := range linha {
			linha[j] = infinito
		}
		matriz = append(matriz, linha)
	}

	// seta os valores de custo dos voos
	for _, voo := range voos {
		matriz[dict[voo.Origem]][dict[voo.Destino]] = voo.Custo
	}

	return dict, matriz, voos
}

// gravaCSV é uma função para a gravação do arquivo CSV para novas
// ligações entre aeroportos ou alteração do valor de custo de
// ligações já existentes
func gravaCSV(voos []Voo) bool {
	var linha string
	retorno := true

	file, err := os.Create(Arquivo)
	checkErr(err)

	defer file.Close()

	linha = ""
	escritor := bufio.NewWriter(file)
	for _, v := range voos {
		linha = fmt.Sprintf("%s,%s,%d", v.Origem, v.Destino, v.Custo)
		fmt.Fprintln(escritor, linha)
	}

	errFlush := escritor.Flush()
	checkErr(errFlush)

	if errFlush != nil {
		retorno = false
	}

	return retorno
}
