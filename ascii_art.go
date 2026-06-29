package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

var emojiError = errors.New("Error: Emoji detected in user input")

func GenerateAsciiArtText(text, bannerType string) (string, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	userText := strings.Split(text, "\n")

	data, err := os.ReadFile("static/assets/banners/" + bannerType)
	if err != nil {
		log.Println("An Error occurred:", err)
		return "", err
	}
	sanitizedData := strings.ReplaceAll(string(data), "\r\n", "\n")
	asciiLines := strings.Split(sanitizedData, "\n")

	result := ""

	for _, word := range userText {
		for i := 1; i <= 8; i++ {
			for _, char := range word {
				if char >= 32 && char <= 126 {
					index := (int(char) - 32)
					result += asciiLines[index*9+i]
				} else {
					log.Printf("Unwanted Character Detected: %q", char)
					return ` /\
 /  \
 / || \
 /  ||  \
 /   ..   \
 /__________\
   E R R O R

 NO EMOJI ALLOWED IN THE TEXT AREA`, emojiError
				}
			}
			result += "\n"
		}
	}

	return result, nil
}
