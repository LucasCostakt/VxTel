package apirest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateHttp(t *testing.T) {
	myStructTestResponse := []struct {
		name string
		want *httpServer
	}{
		{name: "Sucess", want: new(httpServer)},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateHttp()

			AssertResponsebody(t, got, tt.want)
		})
	}
}
func TestNewRoutes(t *testing.T) {
	myStructTestResponse := []struct {
		name string
		want *httpServer
	}{
		{name: "Sucess", want: new(httpServer)},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRoutes(tt.want)
			AssertResponsebody(t, got, tt.want)
		})
	}
}

func TestRequestConsult(t *testing.T) {
	myStructTestResponse := []struct {
		name   string
		url    string
		start  string
		final  string
		client http.Client
		buf    *bytes.Buffer
		want   string
	}{
		{name: "Request Sucess",
			url:    "http://localhost:5000/consult",
			client: http.Client{},
			buf:    bytes.NewBuffer([]byte(`{"start":"018","final":"011","time":200,"plano":"120"}`)),
			want:   `{"start":"018","final":"011","time":"200","value":"167.20","ValueFull":"380.00","plano":"120"}` + "\n",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPost, tt.url, tt.buf)
			request.Header.Set("Content-Type", "application/json")

			response, _ := tt.client.Do(request)
			got, _ := ioutil.ReadAll(response.Body)

			AssertResponsebody(t, string(got), string(tt.want))
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
