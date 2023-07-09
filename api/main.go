package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	r := router.Gerar()

	//fmt.Println(config.SecretKey)

	fmt.Printf("Escutando na porta %d", config.Porta)
	//levanta a aplicação
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
