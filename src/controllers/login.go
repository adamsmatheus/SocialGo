package controllers

import (
	"net/http"
	"webappinterface/src/utils"
)

func CarregaTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "login.html", nil)
}
