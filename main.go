package main

import (
	"Jonath-z/json-parser/utils"
)

func main() {
	tokens := utils.Lexer(`{"hello": "world", "name":"jonathan", "lastName":"zihindula"}`)
	utils.Parser(tokens)
}
