package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type HomePageData struct {
	UserInput      string
	AsciiArtOutput string
	BannerType     string
	EmojiError     string
}

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
		data := HomePageData{
			UserInput:      "",
			AsciiArtOutput: "",
			BannerType:     "standard.txt",
			EmojiError:     "",
		}
		HomePageTemplate.Execute(writer, data)

	case http.MethodPost:
		request.ParseForm()

		bannerType := request.FormValue("banner")
		userInput := request.FormValue("user-input")
		log.Printf("New POST request received: %q\n", request.Form)

		asciiArtText, err := GenerateAsciiArtText(userInput, bannerType)
		log.Printf("New ASCII ART GENERATED:\n")
		fmt.Println(asciiArtText)
		if err != nil {
			if err == emojiError {
				data := HomePageData{
					UserInput:      userInput,
					AsciiArtOutput: "",
					BannerType:     bannerType,
					EmojiError:     asciiArtText,
				}

				HomePageTemplate.Execute(writer, data)
				return
			}
			http.Error(writer, "[HP] 500: An error occured with the specified display style", http.StatusInternalServerError)
			return
		}

		data := HomePageData{
			UserInput:      userInput,
			AsciiArtOutput: asciiArtText,
			BannerType:     bannerType,
			EmojiError:     "",
		}

		HomePageTemplate.Execute(writer, data)
	default:
		http.Error(writer, "[HP] 405: Method Not Allowed", http.StatusMethodNotAllowed)
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
