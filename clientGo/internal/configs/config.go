package configs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	store "github.com/Vortx-Test/clientGo/internal/store"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//faz as transformações necessárias para efetuar o request na api de consulta
func RequestConsult(p store.SendNumbers) store.GetNumbers {
	url := readEnv()
	return requestConsult(p, &http.Client{}, url)
}

//faz as transformações necessárias para efetuar o request na api de consulta
func requestConsult(p store.SendNumbers, client httpClient, url string) store.GetNumbers {
	g := store.GetNumbers{}

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

//Efetua o http resquest passando o metodo a url e os valores a serem enviados
func NewRequest(method string, url string, requestBody []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("Request error NewRequest() ", err)
		return nil, err
	}
	return request, err
}

//Lê a url vinda do Dockerfile
func readEnv() string {
	url := os.Getenv("URL")
	if url == "" {
		url = "err"
	}
	return url
}
