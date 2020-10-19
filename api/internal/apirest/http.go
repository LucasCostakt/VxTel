package apirest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	archives "github.com/Vortx-Test/api/internal/archives"
	caltax "github.com/Vortx-Test/api/internal/caltax"
	store "github.com/Vortx-Test/api/internal/store"
)

type httpServer struct {
	http.Handler
}

func CreateHttp() *httpServer {
	log.Println("Create new httpServer")
	return new(httpServer)
}

//Cria as novas rotas
func NewRoutes(h *httpServer) *httpServer {
	log.Println("Init Routes")
	router := http.NewServeMux()
	//criados os endpoint "/consult"
	router.Handle("/consult", http.HandlerFunc(consult))

	h.Handler = router

	return h
}

//Inicia o server na porta 5000
func StartAPI(routes *httpServer) {
	log.Println("Start API on Port 5000")
	if err := http.ListenAndServe(":5000", routes); err != nil {
		log.Fatal("init server error in StartApi(), ", err)
	}
}

//Handler func que efetua a consulta quando chamado o endpoin /consult
func consult(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cval := &store.ConvertNumbers{}
		bodyNumbers := archives.ReadBodyNumbers(r)

		cval.Time = strconv.Itoa(bodyNumbers.Time)
		cval.StartDDD = bodyNumbers.StartDDD
		cval.FinalDDD = bodyNumbers.FinalDDD
		cval.Plano = bodyNumbers.Plano

		valuesFile := archives.ReadValues()
		cost := archives.ReadFileValue(bodyNumbers.StartDDD, bodyNumbers.FinalDDD, valuesFile)
		caltax.CalculeteTax(bodyNumbers.Time, cost, bodyNumbers.Plano, cval)

		json.NewEncoder(w).Encode(cval)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
