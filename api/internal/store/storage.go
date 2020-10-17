package store

type Numbers struct {
	StartDDD string  `json:"start"`
	FinalDDD string  `json:"final"`
	Value    float32 `json:"value"`
}

type ConvertNumbers struct {
	StartDDD        int     `json:"start"`
	FinalDDD        int     `json:"final"`
	Time            int     `json:"time"`
	Plano           string  `json:"plano"`
	WithFaleMais    float32 `json:"comfalemais"`
	WithoutFaleMais float32 `json:"semfalemais"`
}
