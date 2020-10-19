package archives

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	store "github.com/Vortx-Test/api/internal/store"
)

//Transforma os dados vindos da requisição em um struct onde pode se trabalhar
func ReadBodyNumbers(r *http.Request) store.GetNumbers {
	getn := store.GetNumbers{}
	err := json.NewDecoder(r.Body).Decode(&getn)
	if err != nil {
	}
	return getn
}

//Lê os dados do csv onde está contido os valores das ligações
func ReadValues() []string {
	return readValues("../csv/valores.csv")
}

//Lê os dados do csv onde está contido os valores das ligações
func readValues(fileName string) []string {

	var fileContents []string

	file, err := os.Open(fileName)
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

//Recebe o ddd inicial e final e o conteudo do arquivo de ligações
//Retorna o valor por min gasto que será usado para calcular a ligação
//Caso a ligação entre os dois destinos não exista retorna o valor 0
func ReadFileValue(start, final string, fileContentes []string) float64 {
	for _, dest := range fileContentes {
		s := strings.Split(dest, ",")
		if strings.EqualFold(s[0], start) && strings.EqualFold(s[1], final) {
			i, _ := strconv.ParseFloat(s[2], 32)
			return float64(i)
		}
	}

	return 0
}
