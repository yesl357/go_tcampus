package main

import (
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var urls = []string{
	"http://www.tdogcode.com",
	"http://www.baidu.com",
	"http://www.google.com",
}

func main() {
	for _, url := range urls {
		//每一个url添加一个goroutine,wg+1
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			response, err := http.Get(url)

			if err == nil {
				println(response.Status)
			} else {
				println("error", err.Error())
			}
		}(url)
	}

	wg.Wait()
}
