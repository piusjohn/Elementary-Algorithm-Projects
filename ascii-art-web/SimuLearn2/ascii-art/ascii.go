package asciiart

import (
	"errors"
	"log"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error){
	data, err := os.ReadFile(filename)
	if err != nil{
		log.Fatal(err)
	}
	artfile := strings.ReplaceAll(string(data), "\r\n", "\n")
	splitted := strings.Split(artfile, "\n")

	if len(splitted) == 0 {
		return nil, errors.New("empty file")
	}
	if len(splitted) < 855 {
		return nil, errors.New("corrupt art file")
	}
	asciimap := make(map[rune][]string)
	currentchar := ' '

	for i := 1; i < 855; i += 9{
		if currentchar > 126{
			break
		}
		asciimap[currentchar] = splitted[i : i+8]
		currentchar++
	}
	return asciimap, nil
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

func SplitInput(s string) []string{
	return strings.Split(s, "\\n")
}


func GenerateArt(text string, banner map[rune][]string) string {
	word := SplitInput(text)
	var b strings.Builder
	for i, char := range word{
		if char != ""{
			b.WriteString(strings.Join(RenderLine(char, banner), "\n") + "\n")
		} else if i < len(word)-1{
			b.WriteByte('\n')
		}
	}
	return b.String()
}