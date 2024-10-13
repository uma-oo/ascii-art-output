package ascii

import (
	"log"
	"os"
	"strings"
)

func SplitNewLine(sentence string) []string {
	var slice []string
	var word string

	for i := 0; i < len(sentence); i++ {
		if i < len(sentence)-1 && sentence[i] == '\\' && sentence[i+1] == 'n' {

			slice = append(slice, word)
			word = ""
			i++
		} else {
			word += string(sentence[i])
		}
	}

	slice = append(slice, word)

	if strings.ReplaceAll(sentence, "\\n", "") == "" { // "" {"",""}
		slice = slice[1:]
	}

	return slice
}

func CreateMap(list_of_letters []string) map[string][]string {
	m := make(map[string][]string)
	j := 0
	for i := ' '; i < 127; i++ {
		m[string(i)] = strings.Split(list_of_letters[j], "\n")
		j++

	}
	return m
}

func IsPrintable(word string) bool {
	runes := []rune(word)
	result := true
	for i := 0; i < len(runes); i++ {
		if runes[i] > '~' || runes[i] < ' ' {
			result = false
			break
		} else {
			continue
		}
	}
	return result
}

func Print(words []string, m map[string][]string) string {
	str := ""
	for _, word := range words {
		if !IsPrintable(word) {
			log.Fatal("This sentence contains characters out of the range of printable ascii characters")
		}
		if word == "" {
			str += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for l := 0; l < len(word); l++ {
				str += m[string(word[l])][i]
			}
			str += "\n"
		}
	}
	return str
}

// the fs part

func CheckBanner(banner string) bool {
	banners := []string{"standard", "thinkertoy", "shadow"}
	for _, element := range banners {
		if element == banner {
			return true
		}
	}
	return false
}

func ReadBanner(banner string) []string {
	var file string
	filename, err := os.ReadFile("./banners/" + banner + ".txt")
	CheckError(err)

	if banner != "thinkertoy" {
		file = string(filename)
	} else {
		file = strings.ReplaceAll(string(filename), "\r\n", "\n")
	}
	liste_of_letters := strings.Split(string(file[1:len(file)-1]), "\n\n")

	return liste_of_letters
}

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
