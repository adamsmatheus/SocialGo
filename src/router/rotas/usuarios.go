package rotas

import (
	"net/http"
	"webappinterface/src/controllers"
)

var rotasUsuario = []Rota{
	{
		URI:                "/criar-cadastro",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregaPaginaCadastroUsuario,
		RequerAutenticacao: false,
	},
}
