package webgolang

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:9090",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/tablor", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hey tablor you are on servermux.")
	})

	mux.HandleFunc("/1453", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hey 1453 on constantinenopel.")
	})

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
