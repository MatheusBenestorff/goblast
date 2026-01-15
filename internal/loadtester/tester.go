package loadtester

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func RunLoadTest(url string, totalRequests int) {
	fmt.Printf("Iniciando ataque a %s com %d requests...\n", url, totalRequests)

	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			MakeRequest(url, id)
		}(i) 
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Ataque finalizado! Tempo total: %s\n", elapsed)
}

func MakeRequest(url string, id int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[Erro] Req %d falhou: %v\n", id, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[OK] Req %d -> Status: %d\n", id, resp.StatusCode)
}