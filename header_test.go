package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HeaderHandler(writer http.ResponseWriter, request *http.Request) {
	header := request.Header.Get("content-type")
	fmt.Fprint(writer, header)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	HeaderHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("x-powered-by", "Scheva Pratama")
	fmt.Fprint(writer, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	powered_by := recorder.Header().Get("x-powered-by")
	fmt.Println(powered_by)
}
