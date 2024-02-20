package controllers

import (
	"net/http"
	"webappinterface/src/utils"
)

func CarregaTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "login.html", nil)
}

func CarregaPaginaCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "cadastro.html", nil)
}
