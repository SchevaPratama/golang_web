package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func StatusCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(400) //Bad request
		fmt.Fprint(writer, "Name can't be empty")
	} else {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestStatusCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=zio", nil)
	recorder := httptest.NewRecorder()

	StatusCode(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(result.StatusCode)
	fmt.Println(result.Status)
	fmt.Println(string(body))
}
