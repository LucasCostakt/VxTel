package apirest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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

func consult(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cval := &store.ConvertNumbers{}
		p := readBodyNumbers(r)

		cval.Time = strconv.Itoa(p.Time)
		cval.StartDDD = p.StartDDD
		cval.FinalDDD = p.FinalDDD
		cval.Plano = p.Plano

		valuesFile := readValues()
		cost := readFileValue(p.StartDDD, p.FinalDDD, valuesFile)
		calculeteTax(p.Time, cost, p.Plano, cval)

		json.NewEncoder(w).Encode(cval)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func readBodyNumbers(r *http.Request) store.GetNumbers {
	p := store.GetNumbers{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
	}
	return p
}

func readValues() []string {

	var fileContents []string

	file, err := os.Open("../csv/valores.csv")
	if err != nil {
		log.Println("os Open error in OpenFile() ", err)
		return []string{"os Open error"}
	}
	in := file
	buf := bufio.NewScanner(in)

	for buf.Scan() {
		fileContents = append(fileContents, buf.Text())
	}

	if err := buf.Err(); err != nil {
		log.Println("reading file error in OpenFile() ", err)
		return []string{" reading file error"}
	}

	defer file.Close()

	return fileContents
}

func readFileValue(start, final string, fileContentes []string) float64 {
	for _, dest := range fileContentes {
		s := strings.Split(dest, ",")
		if strings.EqualFold(s[0], start) && strings.EqualFold(s[1], final) {
			i, _ := strconv.ParseFloat(s[2], 32)
			return float64(i)
		}
	}

	return 0
}

func calculeteTax(time int, cost float64, plano string, c *store.ConvertNumbers) error {
	switch plano {
	case "30":

		if time-30 > 0 {
			t := time - 30
			desc := float64(t) * cost
			c.Value = fmt.Sprintf("%.2f", desc*0.10+desc)
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		} else {
			c.Value = "0"
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		}

	case "60":
		if time-60 > 0 {
			t := time - 60
			desc := float64(t) * cost
			c.Value = fmt.Sprintf("%.2f", desc*0.10+desc)
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		} else {
			c.Value = "0"
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		}
	case "120":
		if time-120 > 0 {
			t := time - 120
			desc := float64(t) * cost
			c.Value = fmt.Sprintf("%.2f", desc*0.10+desc)
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		} else {
			c.Value = "0"
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		}
	default:
		return fmt.Errorf("erro")
	}

	return fmt.Errorf("erro")
}
