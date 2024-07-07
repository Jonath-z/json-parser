package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerLeftBrace(t *testing.T) {
	result := Lexer("{")
	expected := Token{
		tokenType: LEFT_BRACE,
		value:     "{",
	}

	assert.Equal(t, result[0].tokenType, expected.tokenType, "Should have LEFT_BRACE as token type")
	assert.Equal(t, result[0].value, expected.value, "Should have { as value")
}

func TestLexerRightBrace(t *testing.T) {
	result := Lexer("}")
	expected := Token{
		tokenType: RIGHT_BRACE,
		value:     "}",
	}
	assert.Equal(t, result[0].tokenType, expected.tokenType, "Should have RIGHT_BRACE as token type")
	assert.Equal(t, result[0].value, expected.value, "Should have } as value")
}

func TestLeftAndRightBrace(t *testing.T) {
	result := Lexer("{}")
	expected := []Token{
		{
			tokenType: LEFT_BRACE,
			value:     "{",
		},
		{
			tokenType: RIGHT_BRACE,
			value:     "}",
		},
	}

	assert.ElementsMatch(t, result, expected, "Should have left and right brance tokens")
}

func TestFullJsonTokens(t *testing.T) {
	result := Lexer(`{"name":"foo" , "value":"bar"}`)
	expected := []Token{
		{
			tokenType: LEFT_BRACE,
			value:     "{",
		},
		{
			tokenType: STRING,
			value:     "name",
		},
		{
			tokenType: COLON,
			value:     ":",
		},
		{
			tokenType: STRING,
			value:     "foo",
		},
		{
			tokenType: COMMA,
			value:     ",",
		},
		{
			tokenType: STRING,
			value:     "value",
		},
		{
			tokenType: COLON,
			value:     ":",
		},
		{
			tokenType: STRING,
			value:     "bar",
		},
		{
			tokenType: RIGHT_BRACE,
			value:     "}",
		},
	}
	assert.ElementsMatch(t, result, expected)
}
