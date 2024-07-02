package utils

import "regexp"

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
	currentPosition := 0
	tokens := []Token{}

	for currentPosition < len(s) {
		character := string(s[currentPosition])

		if character == "{" {
			tokens = append(tokens, Token{
				tokenType: LEFT_BRACE,
				value:     character,
			})
			currentPosition++
		}
		if character == "}" {
			tokens = append(tokens, Token{
				tokenType: RIGHT_BRACE,
				value:     character,
			})
			currentPosition++
		}
		if character == "[" {
			tokens = append(tokens, Token{
				tokenType: LEFT_BRACKET,
				value:     character,
			})
			currentPosition++
		}
		if character == "]" {
			tokens = append(tokens, Token{
				tokenType: RIGHT_BRACKET,
				value:     character,
			})
			currentPosition++
		}
		if character == ":" {
			tokens = append(tokens, Token{
				tokenType: COLON,
				value:     character,
			})
			currentPosition++
		}
		if character == "," {
			tokens = append(tokens, Token{
				tokenType: COMMA,
				value:     character,
			})
			currentPosition++
		}

		if character == "\"" {
			stringSequence := ""
			character = string(s[currentPosition+1])
			for character != "\"" {
				stringSequence += character
				character = string(s[currentPosition+1])
			}
			character = string(s[currentPosition+1])
			tokens = append(tokens, Token{
				tokenType: STRING,
				value:     stringSequence,
			})
			continue
		}

		emptySpace, _ := regexp.Compile(" ")
		if emptySpace.Match([]byte(string(s[currentPosition]))) {
			continue
		}

		numbers, _ := regexp.Compile("[0-9]")
		if numbers.Match([]byte(character)) {
			numberSequence := ""
			for numbers.Match([]byte(character)) {
				numberSequence += character
				character = string(s[currentPosition+1])
			}
			tokens = append(tokens, Token{
				tokenType: NUMBER,
				value:     numberSequence,
			})
			continue
		}

		if character == "t" && string(s[currentPosition+1]) == "r" && string(s[currentPosition+2]) == "u" && string(s[currentPosition+3]) == "e" {
			tokens = append(tokens, Token{
				tokenType: TRUE,
				value:     "true",
			})
			currentPosition += 4
			continue
		}

		if character == "f" && string(s[currentPosition+1]) == "a" && string(s[currentPosition+2]) == "l" && string(s[currentPosition+3]) == "s" && string(s[currentPosition+4]) == "e" {
			tokens = append(tokens, Token{
				tokenType: FALSE,
				value:     "false",
			})
			currentPosition += 5
			continue
		}

		if character == "n" && string(s[currentPosition+1]) == "u" && string(s[currentPosition+2]) == "l" && string(s[currentPosition+3]) == "l" {
			tokens = append(tokens, Token{
				tokenType: NULL,
				value:     "null",
			})
			currentPosition += 4
			continue
		}

		panic("Unknown Character: " + character)
	}

	return tokens
}
