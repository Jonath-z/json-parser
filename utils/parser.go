package utils

import (
	"strings"
)

type ObjectNode struct {
	nodeType string
	value    Token
	children map[string]*ObjectNode
}

func buildNode(tokenType string, value string) *ObjectNode {
	return &ObjectNode{
		nodeType: strings.ToLower(tokenType),
		value: Token{
			tokenType: tokenType,
			value:     value,
		},
		children: make(map[string]*ObjectNode),
	}
}

func parseObject(tokens []Token, currentPosition *int) *ObjectNode {
	*currentPosition++
	objectNode := &ObjectNode{
		nodeType: "object",
		children: make(map[string]*ObjectNode),
	}

	for tokens[*currentPosition].tokenType != RIGHT_BRACE {
		if tokens[*currentPosition].tokenType == COMMA {
			*currentPosition++
		}

		if tokens[*currentPosition].tokenType == STRING {
			key := tokens[*currentPosition].value
			*currentPosition++

			if tokens[*currentPosition].tokenType != COLON {
				panic("Invalid JSON, expected a colon")
			}
			*currentPosition++
			value := parseValue(tokens, currentPosition)
			objectNode.children[key] = value
		} else {
			panic("Invalid JSON")
		}
		*currentPosition++
		if *currentPosition >= len(tokens) {
			break
		}
	}

	return objectNode
}

func parseValue(tokens []Token, currentPosition *int) *ObjectNode {
	currentTokenType := tokens[*currentPosition].tokenType
	switch currentTokenType {
	case STRING, NUMBER, TRUE, FALSE:
		node := buildNode(currentTokenType, tokens[*currentPosition].value)
		*currentPosition++
		return node
	case LEFT_BRACE:
		return parseObject(tokens, currentPosition)
	default:
		panic("Not token to parse")
	}
}

func Parser(tokens []Token) *ObjectNode {
	if len(tokens) == 0 {
		panic("No tokens to parse")
	}

	i := 0
	return parseValue(tokens, &i)
}
