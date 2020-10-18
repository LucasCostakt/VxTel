package store

type GetNumbers struct {
	StartDDD string `json:"start"`
	FinalDDD string `json:"final"`
	Time     int    `json:"time"`
	Plano    string `json:"plano"`
}

type ConvertNumbers struct {
	StartDDD  string `json:"start"`
	FinalDDD  string `json:"final"`
	Value     string `json:"value"`
	Time      string `json:"time"`
	ValueFull string `json"valuefull"`
	Plano     string `json:"plano"`
}
