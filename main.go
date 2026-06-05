package main

import (
	"net/http"
)

func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "templates/index.html")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homePageHandler)

	http.ListenAndServe("5000", mux)
}
