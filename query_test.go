package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("firstName")
	last_name := request.URL.Query().Get("lastName")
	fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
}

func TestParameterQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=zio", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestMultipleParamterQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?firstName=Zio&lastName=Pratama", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleValueParamter(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, ","))
}

func TestMultipleValueParamter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Zio&name=Zordiev&name=Scheva", nil)
	recorder := httptest.NewRecorder()

	MultipleValueParamter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
