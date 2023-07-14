package main

import (
	"fmt"
	"math/rand"
)

func gera_pacote(N, n int) []int {
	pacote := []int{}
	for i := 1; i <= n; i++ {
		figurinha := rand.Intn(N) + 1
		pacote = append(pacote, figurinha)
	}
	return pacote
}

func simula_album(N, n int) int {
	album := map[int]int{}
	pacotes := 0
	for len(album) < N {
		pacote := gera_pacote(N, n)
		for _, v := range pacote {
			album[v]++
		}
		pacotes++
	}
	return pacotes
}

func media(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total / len(x)
}

func main() {
	N := 500  // tamanho do album
	n := 4    // tamanho do pacote
	m := 1000 // quantidade de simulações

	resultados := []int{}
	for i := 1; i <= m; i++ {
		resultados = append(resultados, simula_album(N, n))
	}
	fmt.Println("Média de pacotes a serem comprados:", media(resultados))
}
