package main

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	requestCount := 100

	wg := sync.WaitGroup{}
	wg.Add(requestCount)
	for i := 0; i < requestCount; i++ {
		go func(val int, wg *sync.WaitGroup) {
			client := resty.New()
			resp, _ := client.R().
				Get("http://localhost:3000/")
			logrus.Info(resp)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
