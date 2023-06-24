package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statment, erro := repositorio.db.Prepare("insert into usuarios (nome, nick,email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	resultado, erro := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}

func (repositorio usuarios) Buscar(nomeOunick string) ([]modelos.Usuario, error) {
	nomeOunick = fmt.Sprintf("%%%s%%", nomeOunick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoem from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOunick, nomeOunick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var umUsuario modelos.Usuario

		if erro = linhas.Scan(
			&umUsuario.ID,
			&umUsuario.Nome,
			&umUsuario.Nick,
			&umUsuario.Email,
			&umUsuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, umUsuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoem from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var umUsuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(
			&umUsuario.ID,
			&umUsuario.Nome,
			&umUsuario.Nick,
			&umUsuario.Email,
			&umUsuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}

	}
	return umUsuario, nil
}
