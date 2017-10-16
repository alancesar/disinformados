package utils

import "fmt"
import "../models"

// Printer Imprime o endereço
func Printer(endereco *models.Endereco) {
	fmt.Println("Logradouro ", endereco.Logradouro)
	fmt.Println("Bairro     ", endereco.Bairro)
	fmt.Println("Cidade     ", endereco.Localidade)
	fmt.Println("Estado     ", endereco.UF)
	fmt.Println("IBGE       ", endereco.IBGE)
}
