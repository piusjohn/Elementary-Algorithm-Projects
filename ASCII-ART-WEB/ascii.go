package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, _ := os.ReadFile(filename)
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	splitted := strings.Split(content, "\n")
	hash := make(map[rune][]string)
	if len(splitted) == 0 {
		return nil, errors.New("empty files")
	}
	if len(splitted) < 855 {
		return nil, errors.New("corrupted files")
	}
	currentchar := ' '
	for i := 1; i < len(splitted); i += 9 {
		if currentchar > 126 {
			break
		}
		hash[currentchar] = splitted[i : i+8]
		currentchar++
	}
	return hash, nil
}

func RenderLine(input string, banner map[rune][]string) []string {

	var result []string
	for i := 0; i < 8; i++ {
		var b strings.Builder
		for _, char := range input {
			b.WriteString(banner[char][i])
		}
		result = append(result, b.String())
	}
	return result
}

func SplitInput(text string) []string {
	return strings.Split(text, "\\n")
}

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("unsupported character")
		}
	}
	return 0, nil
}

func GenerateArt(input string, banner map[rune][]string) string {
	r, err := ValidateInput(input)
	if r != 0 && err != nil {
		return ("unsupported character")
	}
	var b strings.Builder
	data := SplitInput(input)
	for i, char := range data {
		if char != "" {
			b.WriteString(strings.Join(RenderLine(char, banner), "\n") + "\n")
		} else if i < len(data)-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("USAGE: go run . 'userinput' standard.txt")
// 		return
// 	}
// 	userinput := os.Args[1]
// 	if userinput == "" {
// 		return
// 	}

// 	bannername := "standard"
// 	if len(os.Args) > 2 {
// 		bannername = os.Args[2]
// 	}
// 	bannerpath := fmt.Sprintf("%s.txt", bannername)

// 	_, err := LoadBanner(bannerpath)
// 	if err != nil {
// 		fmt.Printf("Error: Banner '%s' not found or invalid.\n", bannername)
// 		return
// 	}

// 	data, _ := LoadBanner(bannerpath)

// 	result := GenerateArt(userinput, data)
// 	fmt.Println(result)

// }
