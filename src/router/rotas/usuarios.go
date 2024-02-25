package rotas

import (
	"net/http"
	"webappinterface/src/controllers"
)

var rotasUsuario = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregaPaginaCadastroUsuario,
		RequerAutenticacao: false,
	},
}
