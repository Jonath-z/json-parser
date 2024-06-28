package main

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

func lexer(s string) []Token {
	tokens := []Token{}

	for i := 0; i < len(s); i++ {
		character := string(s[i])

		if character == "{" {
			tokens = append(tokens, Token{
				tokenType: LEFT_BRACE,
				value:     character,
			})
		}
		if character == "}" {
			tokens = append(tokens, Token{
				tokenType: RIGHT_BRACE,
				value:     character,
			})
		}
		if character == "[" {
			tokens = append(tokens, Token{
				tokenType: LEFT_BRACKET,
				value:     character,
			})
		}
		if character == "]" {
			tokens = append(tokens, Token{
				tokenType: RIGHT_BRACKET,
				value:     character,
			})
		}
		if character == ":" {
			tokens = append(tokens, Token{
				tokenType: COLON,
				value:     character,
			})
		}
		if character == "," {
			tokens = append(tokens, Token{
				tokenType: COMMA,
				value:     character,
			})
		}
		if character == "\"" {
			stringSequence := ""
			character = string(s[i+1])
			for character != "\"" {
				stringSequence += character
				character = string(s[i+1])
			}
			character = string(s[i+1])
			tokens = append(tokens, Token{
				tokenType: STRING,
				value:     stringSequence,
			})
			continue
		}

		match, _ := regexp.MatchString("\\s", string(s[i]))
		if match {
			continue
		}

	}

	return tokens
}

func main() {
	fmt.Println("JSON parser")
	lexer("{ }")
}
