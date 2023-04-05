package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//? Função que cria a secret key para login
//! func init (){
//! 	key := make([]byte, 64)

//! 	if _, err := rand.Read(key); err != nil {
//! 		log.Fatal(err)
//! 	}

//! 	fmt.Println(key)
//! 	stringBase64 := base64.StdEncoding.EncodeToString(key)
//! 	fmt.Println(stringBase64)
//! }

func main() {
	config.Load()

	fmt.Printf("Listening and Serving at port %s...\n", config.Port)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r))
}
