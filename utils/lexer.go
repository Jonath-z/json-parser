package utils

import (
	"fmt"
	"regexp"
)

const (
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"
	LEFT_BRACKET  = "LEFT_BRACKET"
	RIGHT_BRACKET = "RIGHT_BRACKET"
	COLON         = "COLON"
	COMMA         = "COMMA"
	STRING        = "STRING"
	NUMBER        = "NUMBER"
	TRUE          = "TRUE"
	FALSE         = "FALSE"
	NULL          = "NULL"
)

type Token struct {
	tokenType string
	value     string
}

func Lexer(s string) []Token {
	fmt.Println(s)
	currentPosition := 0
	tokens := []Token{}

	for currentPosition < len(s) {
		character := string(s[currentPosition])

		switch character {
		case "{":
			tokens = append(tokens, Token{tokenType: LEFT_BRACE, value: character})
			currentPosition++
		case "}":
			tokens = append(tokens, Token{tokenType: RIGHT_BRACE, value: character})
			currentPosition++
		case "[":
			tokens = append(tokens, Token{tokenType: LEFT_BRACKET, value: character})
			currentPosition++
		case "]":
			tokens = append(tokens, Token{tokenType: RIGHT_BRACKET, value: character})
			currentPosition++
		case ":":
			tokens = append(tokens, Token{tokenType: COLON, value: character})
			currentPosition++
		case ",":
			tokens = append(tokens, Token{tokenType: COMMA, value: character})
			currentPosition++
		case `"`:
			stringSequence := ""
			currentPosition++
			for currentPosition < len(s) && string(s[currentPosition]) != `"` {
				stringSequence += string(s[currentPosition])
				currentPosition++
			}
			currentPosition++
			tokens = append(tokens, Token{tokenType: STRING, value: stringSequence})
		default:
			if regexp.MustCompile(`\s`).MatchString(character) {
				currentPosition++
			} else if regexp.MustCompile("[0-9]").MatchString(character) {
				numberSequence := ""
				for currentPosition < len(s) && regexp.MustCompile("[0-9]").MatchString(string(s[currentPosition])) {
					numberSequence += string(s[currentPosition])
					currentPosition++
				}
				tokens = append(tokens, Token{tokenType: NUMBER, value: numberSequence})
			} else if s[currentPosition:currentPosition+4] == "true" {
				tokens = append(tokens, Token{tokenType: TRUE, value: "true"})
				currentPosition += 4
			} else if s[currentPosition:currentPosition+5] == "false" {
				tokens = append(tokens, Token{tokenType: FALSE, value: "false"})
				currentPosition += 5
			} else if s[currentPosition:currentPosition+4] == "null" {
				tokens = append(tokens, Token{tokenType: NULL, value: "null"})
				currentPosition += 4
			} else {
				panic("Unknown character")
			}
		}
	}

	return tokens
}
