package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type TokenType string

const (
	TokenLET         TokenType = "LET"
	TokenEOF         TokenType = "EOF"
	TokenILLEGAL     TokenType = "ILLEGAL" // help us identify the tokens we did no recognize...
	TokenError       TokenType = "ERROR"
	TokenNumber      TokenType = "NUMBER"
	TokenPlus        TokenType = "PLUS"
	TokenMinus       TokenType = "MINUS"
	TokenMultiply    TokenType = "MULTIPLY"
	TokenDivide      TokenType = "DIVIDE"
	TokenLeftParen   TokenType = "LEFT_PAREN"
	TokenRightParen  TokenType = "RIGHT_PAREN"
	TokenIdentifier  TokenType = "IDENTIFIER"
	TokenKeywordIf   TokenType = "IF"
	TokenKeywordElse TokenType = "ELSE"
)

// define the contents of a token...
type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input   string
	current int
}

// function to create a new lexer every time...
func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

/*
Token Generation...
The NextToken method is the core of the lexer. It reads the input string character by character,
skipping whitespace and identifying tokens based on the current character:
*/
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.current >= len(l.input) {
		return Token{Type: TokenEOF}
	}
	ch := l.currentChar()

	switch {
	case unicode.IsDigit(ch):
		return l.scanNumber()
	case ch == '+':
		l.advance()
		return Token{Type: TokenPlus, Value: "+"}
	case ch == '-':
		l.advance()
		return Token{Type: TokenMinus, Value: "-"}
	case ch == '*':
		l.advance()
		return Token{Type: TokenMultiply, Value: "*"}
	case ch == '/':
		l.advance()
		return Token{Type: TokenDivide, Value: "/"}
	case ch == '(':
		l.advance()
		return Token{Type: TokenLeftParen}
	case ch == ')':
		l.advance()
		return Token{Type: TokenRightParen}
	// case unicode.IsLetter(ch):
	// 	return l.scanIdentifier()
	default:
		fmt.Println(Token{Type: TokenILLEGAL})
		os.Exit(0)
	}
	return Token{Type: TokenError}
}

// methpd to return the character at a given index current...
func (l *Lexer) currentChar() rune {
	return rune(l.input[l.current])
}

// method to increment the index of the input..
func (l *Lexer) advance() {
	l.current++
}

// method to skip a white space character...
func (l *Lexer) skipWhitespace() {
	for l.current < len(l.input) && unicode.IsSpace(l.currentChar()) {
		l.advance()
	}
}

// method to handle the numbe tokens...
func (l *Lexer) scanNumber() Token {
	start := l.current
	for l.current < len(l.input) && unicode.IsDigit(l.currentChar()) {
		l.advance()
	}
	return Token{Type: TokenNumber, Value: l.input[start:l.current]}
}

func (l *Lexer) scanIdentifier() Token {
	start := l.current
	for l.current < len(l.input) && (unicode.IsLetter(l.currentChar())) { // || unicode.IsDigit(l.currentChar())) {
		if strings.Fields(l.input)[0] == "let" {
			// l.advance()
			// l.advance()
			// l.advance()
			l.current += 3
			return Token{Type: TokenLET, Value: "LET"}
		}
		l.current += 1
	}
	return Token{Type: TokenIdentifier, Value: l.input[start:l.current]}
}

func main() {
	input := "let 3 + 5 * (10 - 4)"
	lexer := NewLexer(input)

	fmt.Printf("Statement: %s\n\n", input)
	for {
		token := lexer.NextToken()
		if token.Type == TokenEOF {
			break
		}
		fmt.Printf("Token: %s, Value: %s\n", token.Type, token.Value)
	}
}
