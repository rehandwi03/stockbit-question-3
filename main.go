package main

import (
	"log"
	"regexp"
	"strings"
)

func main() {
	log.Println(findFirstStringInBracketRefactor("(rehan)"))
}

func findFirstStringInBracketRefactor(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		indexClosingBracketFound := strings.Index(str, ")")

		if indexFirstBracketFound >= 0 && indexClosingBracketFound >= 0 {
			re := regexp.MustCompile(`\((.*?)\)`)
			submatchall := re.FindAllString(str, -1)
			for _, element := range submatchall {
				element = strings.Trim(element, "(")
				element = strings.Trim(element, ")")

				return element
			}
		} else {
			return ""
		}
	}

	return "string is empty"
}
