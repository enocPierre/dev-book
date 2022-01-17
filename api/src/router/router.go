package router

import "github.com/gorilla/mux"


// Gerar vai retorna un router com as configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}