package modelos

import "time"

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick uint64    `json:"autorNick,omitempty"`
	Curitdas  uint64    `json:"curtidas,omitempty"`
	CriadaEm  time.Time `json:"criadoEm,omitempty"`
}
