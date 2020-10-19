package caltax

import (
	"fmt"
	"reflect"
	"testing"

	store "github.com/Vortx-Test/api/internal/store"
)

func TestReadValues(t *testing.T) {
	myStructTestResponse := []struct {
		name   string
		time   int
		cost   float64
		plano  string
		conNum *store.ConvertNumbers
		want   *store.ConvertNumbers
	}{
		{name: "Calculate tax fale mais 30 Sucess",
			time:   200,
			cost:   float64(1.90),
			plano:  "30",
			conNum: &store.ConvertNumbers{},
			want: &store.ConvertNumbers{
				Value:     "355.30",
				ValueFull: "380.00",
			},
		},
		{name: "Calculate tax fale mais 60 Sucess",
			time:   200,
			cost:   float64(1.90),
			plano:  "60",
			conNum: &store.ConvertNumbers{},
			want: &store.ConvertNumbers{
				Value:     "292.60",
				ValueFull: "380.00",
			},
		},
		{name: "Calculate tax fale mais 120 Sucess",
			time:   200,
			cost:   float64(1.90),
			plano:  "120",
			conNum: &store.ConvertNumbers{},
			want: &store.ConvertNumbers{
				Value:     "167.20",
				ValueFull: "380.00",
			},
		},
		{name: "error",
			time:   200,
			cost:   float64(1.90),
			plano:  "error",
			conNum: &store.ConvertNumbers{},
			want:   &store.ConvertNumbers{},
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {

			CalculeteTax(tt.time, tt.cost, tt.plano, tt.conNum)

			AssertResponsebody(t, tt.conNum, tt.want)
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
