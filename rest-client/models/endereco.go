package models

// Endereco Estrutura do Endereço
type Endereco struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	UF          string
	Unidade     string
	IBGE        string
}
