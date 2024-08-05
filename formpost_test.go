package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {

	// Parsing form and get the value manual
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstname := request.PostForm.Get("first_name")
	lastname := request.PostForm.Get("last_name")

	// Get the value from form directly and parse it automaticly
	// firstName := request.PostFormValue("first_name")
	// lastName := request.PostFormValue("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Zio&last_name=Pratama")
	request := httptest.NewRequest("POST", "http://localhost/", requestBody)
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
