package configs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	store "github.com/Vortx-Test/clientGo/internal/store"
)

func TestRequestConsult(t *testing.T) {
	myStructTestResponse := []struct {
		name   string
		client http.Client
		url    string
		send   store.SendNumbers
		want   store.GetNumbers
	}{
		{name: "Request Consult Sucess",
			client: http.Client{},
			url:    "http://localhost:5000/consult",
			send: store.SendNumbers{
				StartDDD: "018",
				FinalDDD: "011",
				Time:     200,
				Plano:    "120",
			},
			want: store.GetNumbers{
				StartDDD:  "018",
				FinalDDD:  "011",
				Time:      "200",
				Value:     "167.20",
				ValueFull: "380.00",
				Plano:     "120",
			},
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			got := requestConsult(tt.send, &http.Client{}, tt.url)
			AssertResponsebody(t, got, tt.want)
		})
	}
}

func AssertResponsebody(t *testing.T, got, expectedResponse interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expectedResponse) {
		str1 := fmt.Sprintf("%v", got)
		str2 := fmt.Sprintf("%v", expectedResponse)
		t.Errorf("body is wrong, got %q expectedResponse %q\n", str1, str2)
	}
}
