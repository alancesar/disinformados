package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.google.com",
	"http://www.facebook.com",
}

const totalRequests = 10

// HTTPResponse Resposta HTTP
type HTTPResponse struct {
	url      string
	response *http.Response
	err      error
}

func main() {
	results := asyncHTTPRequest(urls)
	for _, result := range results {
		if result != nil && result.response != nil {
			if result.response.StatusCode == http.StatusOK {
				fmt.Println(result.url, "deu bom")
			} else {
				fmt.Println(result.url, "deu ruim")
			}
		}
	}
}

func asyncHTTPRequest(urls []string) []*HTTPResponse {
	ch := make(chan *HTTPResponse)
	responses := []*HTTPResponse{}
	httpClient := http.Client{}

	for _, url := range urls {
		go func(url string) {
			fmt.Println("Fazendo requisição em", url)
			response, err := httpClient.Get(url)
			ch <- &HTTPResponse{url, response, err}

			if err != nil && response != nil && response.StatusCode == http.StatusOK {
				response.Body.Close()
			}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			if r.err != nil {
				fmt.Println("Falha na requisição de", r.url)
			}

			responses = append(responses, r)

			if len(responses) == len(urls) {
				fmt.Println()
				return responses
			}

		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}
