package rotas

import (
	"api/src/router/controller"
	"net/http"
)

var rotasUsuarios = []Rota {
	{
		URI: "/usarios",
		Metodo : http.MethodPost,
		Funcao:  controller.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usarios",
		Metodo : http.MethodGet,
		Funcao:  controller.BuscarUsauarios,
		RequerAutenticacao: false,

	},
	{
		URI: "/usarios/{usuarioId}",
		Metodo : http.MethodGet,
		Funcao:  controller.BuscarUsuario,
		RequerAutenticacao: false,

	},
	{
		URI: "/usarios/{usaurioId}",
		Metodo : http.MethodPut,
		Funcao:  controller.AtualizandoUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usarios/{usaurioId}",
		Metodo : http.MethodDelete,
		Funcao:  controller.DeletandoUsuario,
		RequerAutenticacao: false,
	},
}