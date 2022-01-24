package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

// Preparar vai chamar os methodos para validar e formatar o usuario recebido
func (usuario *Usuario) validar (etapa string) error {

	if usuario.Nome == "" {
		return errors.New("O Nome é obrigatorio e nao pode ser em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatorio e nao pode ser em branco")
	}

	if usuario.Email == "" {
		return errors.New("O Email é obrigatorio e nao pode ser em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email inserido é invalido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha  é obrigatorio e nao pode ser em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}