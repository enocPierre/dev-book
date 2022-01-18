package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retorna un router com as config
func Gerar() *mux.Router {
	r :=  mux.NewRouter()
	return rotas.Configurar(r)
}