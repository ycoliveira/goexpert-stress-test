package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	TotalTime          time.Duration
	TotalRequests      int
	SuccessfulRequests int
	StatusCodes        map[int]int
}

func main() {
	var url string
	var requests int
	var concurrency int

	flag.StringVar(&url, "url", "", "URL do serviço a ser testado.")
	flag.IntVar(&requests, "requests", 0, "Número total de chamadas.")
	flag.IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas.")
	flag.Parse()

	if url == "" || requests <= 0 || concurrency <= 0 {
		fmt.Println("Por favor, forneça a URL do serviço, o número total de chamadas e a quantidade de chamadas simultâneas.")
		return
	}

	log.Println("Iniciando teste de stress...")

	results := make(chan Result)
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := makeRequest(url, requests/concurrency)
			results <- result
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalRequests := 0
	successfulRequests := 0
	statusCodes := make(map[int]int)

	for result := range results {
		totalRequests += result.TotalRequests
		successfulRequests += result.SuccessfulRequests
		for status, count := range result.StatusCodes {
			statusCodes[status] += count
		}
	}

	totalDuration := time.Since(start)

	fmt.Println("Relatório de Teste:")
	fmt.Printf("Tempo total gasto na execução: %v\n", totalDuration)
	fmt.Printf("Quantidade total de chamadas realizadas: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successfulRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for status, count := range statusCodes {
		fmt.Printf("Status %d: %d\n", status, count)
	}
}

func makeRequest(url string, requests int) Result {
	var result Result

	start := time.Now()
	for i := 0; i < requests; i++ {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("Erro ao realizar request: %v\n", err)
		}
		defer resp.Body.Close()

		result.TotalRequests++
		if resp.StatusCode == http.StatusOK {
			result.SuccessfulRequests++
		} else {
			if result.StatusCodes == nil {
				result.StatusCodes = make(map[int]int)
			}
			result.StatusCodes[resp.StatusCode]++
		}
	}
	result.TotalTime = time.Since(start)

	return result
}
