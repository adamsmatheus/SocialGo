package router

import (
	"github.com/gorilla/mux"
	"webappinterface/src/router/rotas"
)

func Gerar() *mux.Router {

	r := mux.NewRouter()
	return rotas.Configurar(r)
}
