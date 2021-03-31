package scanner

import "fmt"

type Token struct {
	tokentype TokenType
	lexeme    TokenType
	literal   interface{}
	line      int
}

func (token *Token) toString() string {
	return fmt.Sprintf("%d %v %s", token.tokentype, token.lexeme, token.literal)
}

type Scanner struct {
	source []byte
}

func (scanner *Scanner) readSource(source []byte) {
	scanner.source = source
}

func (scanner *Scanner) error(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where string, msg string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, msg)
}

type TokenType int

const (
	// Single character
	LEFT_PAREN  TokenType = iota
	RIGHT_PAREN TokenType = iota
	LEFT_BRACE  TokenType = iota
	RIGHT_BRACE TokenType = iota
	COMMA       TokenType = iota
	DOT         TokenType = iota
	MINUS       TokenType = iota
	PLUS        TokenType = iota
	SEMICOLON   TokenType = iota
	SLASH       TokenType = iota
	STAR        TokenType = iota

	// One or two character tokens
	BANG          TokenType = iota
	BANG_EQUAL    TokenType = iota
	EQUAL         TokenType = iota
	EQUAL_EQUAL   TokenType = iota
	GREATER       TokenType = iota
	GREATER_EQUAL TokenType = iota
	LESS_EQUAL    TokenType = iota

	IDENTIFIER TokenType = iota
	STRING     TokenType = iota
	NUMBER     TokenType = iota

	AND    TokenType = iota
	CLASS  TokenType = iota
	ELSE   TokenType = iota
	FALSE  TokenType = iota
	FUN    TokenType = iota
	FOR    TokenType = iota
	IF     TokenType = iota
	NIL    TokenType = iota
	OR     TokenType = iota
	PRINT  TokenType = iota
	RETURN TokenType = iota
	SUPER  TokenType = iota
	THIS   TokenType = iota
	TRUE   TokenType = iota
	VAR    TokenType = iota
	WHILE  TokenType = iota

	EOF TokenType = iota
)
