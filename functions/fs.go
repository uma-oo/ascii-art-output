package ascii

import "log"

// make the standard banner the default one
	// ila mat3tash l arg standard will be used

func Fs(sentence string, banner ...string) string {
	banner_used := "standard"
	if len(banner) == 1 {
		banner_used = banner[0]
	} else if len(banner) > 1 {
		log.Fatal("You can only choose one banner!!")
	}

	if !CheckBanner(banner_used) {
		log.Fatal("\nUsage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
 
	liste_of_letters := ReadBanner(banner_used)

	m := CreateMap(liste_of_letters)
	words := SplitNewLine(sentence)

	return Print(words, m)
}
