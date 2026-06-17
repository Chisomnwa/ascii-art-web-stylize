package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	HomePageTemplate, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("TEMPLATE ERROR:", err)
		http.Error(writer, "500: Template error occured", http.StatusInternalServerError)
		return
	}

	if request.URL.Path != "/" {
		// http.NotFound(writer, request)
		http.Error(writer, "[HP] 404: File Not Found", http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodGet:
		HomePageTemplate.Execute(writer, "")
	case http.MethodPost:
		request.ParseForm()

		bannerType := request.FormValue("banner")
		userInput := request.FormValue("user-input")

		asciiArtText, err := GenerateAsciiArtText(userInput, bannerType)
		if err != nil {
			http.Error(writer, "[HP] 500: An error occured with the specified text style", http.StatusInternalServerError)
			return
		}

		HomePageTemplate.Execute(writer, asciiArtText)
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
