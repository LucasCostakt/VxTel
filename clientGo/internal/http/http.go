package http

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	configs "github.com/Vortx-Test/clientGo/internal/configs"
	store "github.com/Vortx-Test/clientGo/internal/store"
)

type httpServer struct {
	http.Handler
}

var templates *template.Template

func CreateHttp() *httpServer {
	log.Println("Create new httpServer")
	return new(httpServer)
}

func CreateTemplates() {
	log.Println("Create html Templates")
	templates = template.Must(template.ParseGlob("../internal/http/templates/*.html"))
}

//Cria as novas rotas
func NewRoutes(h *httpServer) *httpServer {
	log.Println("Init Routes")
	router := http.NewServeMux()
	//criados os endpoint "/" e "/consult"
	router.Handle("/", http.HandlerFunc(pageConsult))

	h.Handler = router

	return h
}

//Inicia a interface web na porta 5050
func StartAPI(routes *httpServer) {
	log.Println("Start API on Port 5050")
	if err := http.ListenAndServe(":5050", routes); err != nil {
		log.Fatal("init server error in StartApi(), ", err)
	}
}

func pageConsult(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//caso o metodo for Get vai exibir o template html sem o conteudo de consulta
	case http.MethodGet:
		err := templates.ExecuteTemplate(w, "consult.html", nil)
		if err != nil {
			log.Println("Cannot Get View pageConsult()", err)
		}
		//caso o metodo for post vai exibir o template html com o conteudo de consulta
	case http.MethodPost:
		p := store.SendNumbers{}

		//coleta os valores do html
		p.StartDDD = r.FormValue("start")
		p.FinalDDD = r.FormValue("final")
		p.Time, _ = strconv.Atoi(r.FormValue("time"))
		p.Plano = r.FormValue("plano")

		//faz a consulta
		response := configs.RequestConsult(p)

		//executa o template passando a melhor rota
		err := templates.ExecuteTemplate(w, "consult2.html", response)
		if err != nil {
			log.Println("Cannot Get View pageConsult()", err)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
