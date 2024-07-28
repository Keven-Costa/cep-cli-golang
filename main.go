package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {

	saida := 0
	var entradaUsuario int
	var cep string 
	var url string = "viacep.com.br/ws/01001000/json/"

	for saida != 1 {

		fmt.Println("Digite: \n1 - Para sair \n2 - Para continuar")

		fmt.Scanln(&entradaUsuario)

		if entradaUsuario == 1 {
			break
		}

		fmt.Print("Digite seu CEP:")
		fmt.Scanln(&cep)
		
		url = "https://viacep.com.br/ws/" + cep + "/json/"

		//Fazendo Requisição
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Erro ao fazer a requisição: %v", err)
		}
		defer resp.Body.Close()

		//Lendo o corpo da resposta
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Erro ao ler o corpo da resposta: %v", err)
		}

		fmt.Println("Resposta da API:", string(body))
		
	}


    
}
