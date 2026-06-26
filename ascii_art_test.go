package main

import (
	"os"
	"testing"
)

func Test_With_Valid_Text_And_Missen_Banner(tester *testing.T) {
	sampleText := "Chisom"
	sampleBannerType := "random.txt"

	result, err := GenerateAsciiArtText(sampleText, sampleBannerType)
	if err == nil {
		tester.Fatalf("Expected an error %v, but got %v", os.ErrNotExist, nil)
	}

	if result != "" {
		tester.Fatalf("Expected \"\" but got %v", result)
	}
}

func Test_With_Valid_Text_And_Valid_Banner(tester *testing.T) {
	text := "1Hello 2There"
	bannerType := "standard.txt"
	result, err := GenerateAsciiArtText(text, bannerType)
	expectedResult := `     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`
	if result != expectedResult {
		tester.Fatalf("Expected %v\n but got %v\n", expectedResult, result)
	}

	if err != nil {
		tester.Fatalf("Did not expect any error, but got %v", err)
	}
}
