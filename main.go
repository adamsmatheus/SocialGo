package main

import (
	"log"
	"net/http"
	"webappinterface/src/router"
	"webappinterface/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
