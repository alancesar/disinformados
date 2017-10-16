package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./models"
	"./utils"
)

func main() {
	var endereco models.Endereco
	endereco = models.Endereco{}

	bytes := get("13172150")

	json.Unmarshal(bytes, &endereco)

	utils.Printer(&endereco)
}

func get(cep string) []byte {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro ao acessar", url)
	}

	defer response.Body.Close()

	bytes, _ := ioutil.ReadAll(response.Body)
	return bytes
}
