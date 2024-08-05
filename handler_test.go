package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// // Use the writer functions
		// w.Write([]byte("Hello world"))

		// Use the formatter from fmt package
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	fmt.Println("Server running at: http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Home Page")
	})
	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "About Page")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Image Page")
	})
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Image Thumbnail Page")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Println("Server running at: http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	fmt.Println("Server running at: http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
