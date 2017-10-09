package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// TODO
// Compreender o porquê de passar o ponteiro do struct ao invés de sua referência

// HTTPResponse Resposta HTTP
type HTTPResponse struct {
	url      string
	response *http.Response
	err      error
}

const requests = 100

var urls = []string{
	"http://localhost:3000",
	"http://localhost:3001",
}

func main() {
	init := time.Now().UnixNano() / int64(time.Millisecond)
	results := asyncHTTPRequest(requests, urls)

	for _, result := range results {
		if result != nil && result.response != nil {
			if result.err == nil && result.response.StatusCode == http.StatusOK {
				bytes, _ := ioutil.ReadAll(result.response.Body)
				fmt.Println(string(bytes))
			}
		}
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("Tempo total ->", (end - init), "milissegundos")
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
	response, err := http.Get(url)
	channel <- &HTTPResponse{url, response, err}
}
