package archives

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadValues(t *testing.T) {
	myStructTestResponse := []struct {
		name     string
		fileName string
		want     []string
	}{
		{name: "read values Sucess",
			fileName: "../../csv/teste.csv",
			want:     []string{"011,016,1.90"},
		},
		{name: "read values error",
			fileName: "",
			want:     []string{"os Open error"},
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			got := readValues(tt.fileName)

			AssertResponsebody(t, got, tt.want)
		})
	}
}

func TestReadFileValue(t *testing.T) {
	myStructTestResponse := []struct {
		name          string
		start         string
		final         string
		fileContentes []string
		want          string
	}{
		{name: "read values Sucess",
			start:         "011",
			final:         "016",
			fileContentes: []string{"011,016,1.90"},
			want:          "1.90",
		},
		{name: "read values error",
			start:         "012",
			final:         "016",
			fileContentes: []string{"011,016,1.90"},
			want:          "0.00",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			i := ReadFileValue(tt.start, tt.final, tt.fileContentes)
			got := fmt.Sprintf("%.2f", i)
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
