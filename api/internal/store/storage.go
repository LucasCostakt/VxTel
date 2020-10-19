package store

//usado quando recebe os parametros vindos da requisição http
//utilizado para fazer os calculos
type GetNumbers struct {
	StartDDD string `json:"start"` // DDD de partida
	FinalDDD string `json:"final"` // DDD de destino
	Time     int    `json:"time"`  // Tempo da ligação
	Plano    string `json:"plano"` // Qual o plano fale mais a ser consultado
}

//Retorna esses valores com os calculos já efetuados
type ConvertNumbers struct {
	StartDDD  string `json:"start"`    // DDD de partida
	FinalDDD  string `json:"final"`    // DDD de destino
	Time      string `json:"time"`     // Tempo da ligação
	Value     string `json:"value"`    // Valor da Ligação com o plano fale  mais
	ValueFull string `json"valuefull"` // Valor da ligação sem o plano fale mais
	Plano     string `json:"plano"`    // Qual o plano fale mais que foi consultado
}
