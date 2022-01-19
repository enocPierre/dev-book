package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// usuarios representa um repositorio de usuarios
type Usuárioios struct {
	db *sql.DB
}

//novoRepositorioDeUsuarios cria um repositorio de usaurio
func NovoRepositoriosDeUsuarios(db *sql.DB) *Usuárioios {
	return &Usuárioios{db}
}

func (repositorio Usuárioios) Criar(usuario modelos.Usuárioio) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()
	
	reusltado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := reusltado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}