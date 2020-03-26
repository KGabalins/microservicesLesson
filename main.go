package main

import (
	"net/http"

	"github.com/KGabalins/project1/handlers"
)

func main() {

	hh := handlers.NewHello()

	http.ListenAndServe(":9090", nil)
}
