package domain

type Deputado struct {
	ID          int64  `json:"id"`
	Nome        string `json:"nome"`
	Partido     string `json:"partido"`
	NumeroVotos int64  `json:"numero_votos"`
}
