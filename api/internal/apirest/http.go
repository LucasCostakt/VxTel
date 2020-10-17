package apirest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Vortx-Test/api/internal/store"
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
	//criados os endpoint "/" e "/consult"
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

func consult(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p := store.Numbers{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "start: "+p.StartDDD+"\n"+"final: "+p.FinalDDD)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// func readValues(start, final string) s.Numbers {
// 	return nil
// }

// func convertValues(n s.Numbers) []s.ConvertNumbers {
// 	return
// }

// func convertJson(n []s.ConvertNumbers) []byte {
// 	return
// }
