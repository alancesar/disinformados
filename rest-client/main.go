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

	bytes, _ := get("13086902")

	fmt.Println("JSON:\n", string(bytes))
	fmt.Println("-------------------------------------------------")

	json.Unmarshal(bytes, &endereco)

	utils.Printer(&endereco)
}

// {
//     "cep": "13086-902",
//     "logradouro": "Rua Doutor Ricardo Benetton Martins",
//     "complemento": "s/n",
//     "bairro": "Polo II de Alta Tecnologia (Campinas)",
//     "localidade": "Campinas",
//     "uf": "SP",
//     "unidade": "",
//     "ibge": "3509502",
//     "gia": "2446"
// }
func get(cep string) (bytes []byte, err error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	// url := "http://localhost:3000"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro ao acessar", url)
	}

	defer response.Body.Close()

	bytes, err = ioutil.ReadAll(response.Body)
	return
}
