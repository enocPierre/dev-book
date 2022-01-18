package rotas

import (
	"api/src/router/controller"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsauario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usaurioId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizandoUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usaurioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletandoUsuario,
		RequerAutenticacao: false,
	},
}
