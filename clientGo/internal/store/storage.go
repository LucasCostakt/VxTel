package store

type SendNumbers struct {
	StartDDD string `json:"start"`
	FinalDDD string `json:"final"`
	Time     int    `json:"time"`
	Plano    string `json:"plano"`
}

type GetNumbers struct {
	StartDDD  string `json:"start"`
	FinalDDD  string `json:"final"`
	Time      string `json:"time"`
	Value     string `json:"value"`
	ValueFull string `json"valuefull"`
	Plano     string `json:"plano"`
}
