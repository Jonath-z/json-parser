package utils

import "strings"

type ObjectNode struct {
	nodeType string
	value    Token
}

func buildNode(tokenType string, value string) *ObjectNode {
	node := &ObjectNode{}
	node.nodeType = strings.ToLower(tokenType)
	node.value.tokenType = tokenType
	node.value.value = value
	return node
}

func parseValue(tokens []Token, currentPosition int) *ObjectNode {
	currentTokenType := tokens[currentPosition].tokenType
	switch currentTokenType {
	case STRING:
		return buildNode(currentTokenType, tokens[currentPosition].value)
	case NUMBER:
		return buildNode(currentTokenType, tokens[currentPosition].value)
	case TRUE:
		return buildNode(currentTokenType, tokens[currentPosition].value)
	case FALSE:
		return buildNode(currentTokenType, tokens[currentPosition].value)
	case LEFT_BRACE:
		parseObject := func() *ObjectNode {
			currentToken := tokens[currentPosition+1]
			objectNode := &ObjectNode{}

			for currentToken.tokenType != RIGHT_BRACE {
				if currentToken.tokenType == STRING {
					key := currentToken.value
					nextToken := tokens[currentPosition+1]
					if nextToken.tokenType != COLON {
						panic("Invalid JSON, expected a colon")
					}
					value := parseValue(tokens, currentPosition+1)
					objectNode.nodeType = key
					objectNode.value = Token{
						tokenType: value.value.tokenType,
						value:     value.value.value,
					}

				} else {
					panic("Invalid JSON, Expected String key in object")
				}
				nextToken := tokens[currentPosition+1]

				if nextToken.tokenType != COMMA {
					panic("Invalid JSON, Expected a comma")
				}
			}

			return objectNode
		}
		return parseObject()
	default:
		panic("Not token to parse")
	}
}

func Parser(tokens []Token) {
	if len(tokens) == 0 {
		panic("Not token to parse")
	}

	i := 0
	parseValue(tokens, i)
}
