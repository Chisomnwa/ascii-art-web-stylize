package main

import (
	"log"
	"os"
	"strings"
)

func GenerateAsciiArtText(text, bannerType string) (string, error) {
	userText := strings.Split(text, "\n")

	data, err := os.ReadFile("static/assets/banners/" + bannerType)
	if err != nil {
		log.Println("An Error occurred:", err)
		return "", err
	}
	asciiLines := strings.Split(string(data), "\n")

	result := ""

	for _, word := range userText {
		for i := 1; i <= 8; i++ {
			for _, char := range word {
				if char >= 32 && char <= 126 {
					index := (int(char) - 32)
					result += asciiLines[index*9+i]
				}
			}
			result += "\n"
		}
	}

	return result, nil
}
