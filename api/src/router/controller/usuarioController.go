package controller

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuárioio!"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos Usuárioio!"))
}

func BuscarUsauario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuárioio por ID!"))
}

func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuárioio!"))
}

func DeletandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuárioio!"))
}