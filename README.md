# cep-cli-golang

### Descrição
Este projeto é uma CLI (Command Line Interface) em Golang que consome a API ViaCep para buscar informações de endereço a partir de um CEP fornecido pelo usuário.

### Funcionalidades
- Consultar informações de endereço usando um CEP

- Interface interativa na linha de comando

### Instalação
Passos para instalar e rodar o projeto localmente:


```
# Clone o repositório
git clone https://github.com/Keven-Costa/cep-cli-golang.git

# Execute o projeto
go run main.go
```

### Uso
Exemplos de como usar o projeto:

```
Digite: 
1 - Para sair 
2 - Para continuar
2
Digite seu CEP: 01001000
Resposta da API: {
  "cep": "01001-000",
  "logradouro": "Praça da Sé",
  "complemento": "lado ímpar",
  "bairro": "Sé",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}
Digite: 
1 - Para sair 
2 - Para continuar
1



```