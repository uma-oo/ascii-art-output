package ascii

import (
	"errors"
	"log"
	"os"
	"strings"
)

func CheckFlagFormat(flag string) bool {
	return strings.HasPrefix(flag, "--output=")
}

func ParseFlag(argument string) string {
	if !CheckFlagFormat(argument) || !strings.HasSuffix(argument, ".txt") {
		log.Fatal("\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standar")
	}

	return argument[len("--output="):]
}

func CreateFile(name_file string, directory_path string) (*os.File, error) {
	// Create the directory if it does not exist
	// mkdirall here is used to handle if we give dirs/dir and create them
	if err := os.MkdirAll(directory_path, os.ModePerm); err != nil {
		return nil, errors.New("directory not found")
	}
	filepath := directory_path + "/" + name_file
	newfile, err := os.Create(filepath)
	if err != nil {
		return nil, errors.New("failed to create the file")
	}

	return newfile, nil
}

func Output(filename string, sentence string, banner ...string) {
	banner_used := "standard"
	if len(banner) == 1 {
		banner_used = banner[0]
	} else if len(banner) > 1 {
		log.Fatal("You can only choose one banner!!")
	}
	// check if the banner is a valid type and there is no typo
	if !CheckBanner(banner_used) {
		log.Fatal("\nUsage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
	list_of_letters := ReadBanner(banner_used)
	m := CreateMap(list_of_letters)
	words := SplitNewLine(sentence)
	result := Print(words, m)
	file, err_creation := CreateFile(filename, "result")
	CheckError(err_creation)
	defer file.Close()
	_, err := file.WriteString(result)
	CheckError(err)
	file.Sync()
}
