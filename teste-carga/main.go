package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPResponse Resposta HTTP
type HTTPResponse struct {
	url      string
	response *http.Response
	err      error
}

const requests = 30000

var urls = []string{"http://localhost:3000"}

func main() {
	init := time.Now().UnixNano() / int64(time.Millisecond)
	results := asyncHTTPRequest(requests, urls)
	sucess := 0

	for _, result := range results {
		if result != nil && result.response != nil {
			if result.err == nil && result.response.StatusCode == http.StatusOK {
				sucess++
				bytes, _ := ioutil.ReadAll(result.response.Body)
				fmt.Println(string(bytes))
			}
		}
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("Tempo total ->", (end - init), "milissegundos")
	fmt.Println(requests, "requests realizados")
	fmt.Println((requests - sucess), "deram ruim")
}

func asyncHTTPRequest(requests int, urls []string) []*HTTPResponse {
	channel := make(chan *HTTPResponse)
	responses := []*HTTPResponse{}
	totalRequests := len(urls) * requests

	for i := 0; i < requests; i++ {
		for _, url := range urls {
			go ping(url, channel)
		}
	}

	for {
		select {
		case response := <-channel:
			if response.err != nil {
				fmt.Println("Falha na requisição de", response.url)
				fmt.Println(response.err)
			}

			responses = append(responses, response)

			if len(responses) == totalRequests {
				return responses
			}
		}
	}
}

func ping(url string, channel chan *HTTPResponse) {
	fmt.Println("Fazendo requisição em", url)
	client := http.Client{}
	response, err := client.Get(url)

	if response != nil {
		defer response.Body.Close()
	}

	channel <- &HTTPResponse{url, response, err}
}
