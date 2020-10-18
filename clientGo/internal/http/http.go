package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	store "github.com/Vortx-Test/clientGo/internal/store"
)

type httpServer struct {
	http.Handler
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
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
		response := RequestConsult(p)

		//executa o template passando a melhor rota
		err := templates.ExecuteTemplate(w, "consult2.html", response)
		if err != nil {
			log.Println("Cannot Get View pageConsult()", err)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func RequestConsult(p store.SendNumbers) store.GetNumbers {
	return requestConsult(p, &http.Client{})
}

func requestConsult(p store.SendNumbers, client httpClient) store.GetNumbers {
	g := store.GetNumbers{}

	url := readEnv()

	js, err := json.Marshal(p)
	if err != nil {
		log.Println("Json marshal error in requestConsult(), ", err)
	}

	request, err := NewRequest(http.MethodPost, url, js)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("Request error in requestConsult(), ", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("Response error in requestConsult(), ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Ioutil ReadAll error in requestConsult(), ", err)
	}

	err = json.Unmarshal(body, &g)
	if err != nil {
		log.Println("Json Unmarshal error in requestConsult(), ", err)
	}

	return g
}

func NewRequest(method string, url string, requestBody []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("Request error NewRequest() ", err)
		return nil, err
	}
	return request, err
}

func readEnv() string {
	url := os.Getenv("URL")
	if url == "" {
		url = "err"
	}
	return url
}
