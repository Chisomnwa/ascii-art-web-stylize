package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		// http.NotFound(writer, request)
		http.Error(writer, "[HP] 404: File Not Found", http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodGet:
		http.ServeFile(writer, request, "templates/index.html")
	case http.MethodPost:
		// Process form here
		request.ParseForm()

		fmt.Println("NEW FORM RECEIVED")
		fmt.Println(request.Form)
		http.ServeFile(writer, request, "templates/index.html")
	default:
		http.Error(writer, "405: Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./static"))

	mux.HandleFunc("/", homePageHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("Server is running on port: http://localhost:5500")
	err := http.ListenAndServe(":5500", mux)
	if err != nil {
		log.Fatal(err)
	}
}
