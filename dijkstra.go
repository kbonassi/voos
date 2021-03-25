package main

// infinito é uma constante com valor alto, simbolizando um valor para infinito
const infinito = 4294967295

// null é uma constante para definir um valor negativo que será utilizado como nulo
const null = -1

// initialize_single_source é uma função que inicializa os vetores
// de distancias e predecessores
func initialize_single_source(grafo [][]int, inicio int) ([]int, []int) {
	n := len(grafo[0]) // linha 0 apenas para pegar a qtd de vertices
	dist := make([]int, n, 10)
	pred := make([]int, n, 10)
	for i := 0; i < n; i++ {
		dist[i] = infinito
		pred[i] = null
	}
	dist[inicio] = 0
	return dist, pred
}

// extract_min é uma função cujo objetivo é extrair
// a menor distância entre os vertices adjacentes
// ao vertice de pesquisa e que estejam abertos
func extract_min(vertices []int, fechados []int) int {
	min := null
	for v := range vertices {
		if fechados[v] == null {
			if min == null {
				min = v
			} else if vertices[v] < vertices[min] {
				min = v
			}
		}
	}
	return min
}

// Dijkstra é uma função que implementa o algorítmo de Dijkstra
// cuja função é buscar um caminho mais econômico entre dois
// pontos (origem-destino). A implementação deste algorítmo foi
// baseada no algorítmo de Dijkstra descrito no livro
// "Algoritmos, teoria e prática" de Cormen, Thomas H. et all.
func Dijkstra(grafo [][]int, inicio int, fim int) ([]int, int) {
	dist, pred := initialize_single_source(grafo, inicio)
	var u int
	n := len(dist)
	S := make([]int, n, 10)
	for v := range S { // nula todos os visitados
		S[v] = null
	}
	Q := dist
	for i := 0; i < n; i++ {
		Q = dist
		u = extract_min(Q, S)
		S[u] = 1
		for v := 0; v < n; v++ { // cada vertice existente em grafo que sai de u
			if dist[v] > dist[u]+grafo[u][v] {
				dist[v] = dist[u] + grafo[u][v]
				pred[v] = u
			}
		}
	}
	paramRota := make([]int, 0, 10)
	rota := print_path(pred, inicio, fim, paramRota)
	custo := 0
	if len(rota) > 1 {
		custo = dist[fim]
	}
	return rota, custo
}

// print_path é uma função que retorna o caminho determinado pelo
// algorítmo de Dijkstra em ordem de origem-destino
func print_path(pred []int, s int, v int, rota []int) []int {
	if v == null {
		vazio := make([]int, 0)
		return vazio
	}

	if v == s {
		rota = append(rota, s)
		return rota
	} else {
		rota = print_path(pred, s, pred[v], rota)
		rota = append(rota, v)
	}

	return rota
}
