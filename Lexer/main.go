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
	TokenSemiColon   TokenType = "SEMICOLON"
	TokenFunction    TokenType = "FUNCTION"
	TokenComment     TokenType = "COMMENT"
	TokenEqual       TokenType = "EQUAL"
	TokenLessThan    TokenType = "LESS_THAN"
	TokenGreaterThan TokenType = "GREATER_THAN"
	TokenColon       TokenType = "COLON"
	TokenComma       TokenType = "COMMA"
	TokenLeftBrace   TokenType = "BRACE"
	TokenRightBrace  TokenType = "RIGHT_BRACE"
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

	case ch == '=':
		l.advance()
		return Token{Type: TokenEqual, Value: "="}
	case ch == '<':
		l.advance()
		return Token{Type: TokenLessThan, Value: "<"}
	case ch == '>':
		l.advance()
		return Token{Type: TokenGreaterThan, Value: ">"}
	case ch == ':':
		l.advance()
		return Token{Type: TokenColon, Value: ":"}
	case ch == ',':
		l.advance()
		return Token{Type: TokenComma, Value: ","}
	case ch == '{':
		l.advance()
		return Token{Type: TokenLeftBrace, Value: "{"}
	case ch == '}':
		l.advance()
		return Token{Type: TokenRightBrace, Value: "}"}
	case ch == ';':
		l.advance()
		return Token{Type: TokenSemiColon, Value: ";"}
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
		return l.scanComment()
	case ch == '(':
		l.advance()
		return Token{Type: TokenLeftParen}
	case ch == ')':
		l.advance()
		return Token{Type: TokenRightParen}
	case unicode.IsLetter(ch):
		return l.scanIdentifier()
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

// scanComment scans comments or division tokens
func (l *Lexer) scanComment() Token {
	if l.current < len(l.input) {
		var count int
		for _, c := range l.input {
			if c == '/' {
				count++
			} else {
				continue
			}
		}
		// println(count)
		if count == 1 {
			l.advance() // Move past the single '/'
			return Token{Type: TokenDivide, Value: "/"}
		} else if count > 1 {
			start := l.current
			for l.current < len(l.input) && l.input[l.current] != '\n' {
				l.advance()
			}
			return Token{Type: TokenComment, Value: l.input[start:l.current]}
		}
	}
	// If no comment or divide token is found, return an illegal token
	return Token{Type: TokenILLEGAL}
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
		} else if strings.Fields(l.input)[0] == "Func" || strings.Fields(l.input)[0] == "func" {
			l.current += 4
			return Token{Type: TokenFunction, Value: "FUNCTION"}
		}
		l.current += 1
	}
	return Token{Type: TokenIdentifier, Value: l.input[start:l.current]}
}

func main() {
	input := "//let 3 / 5 *  (10 - 4)"
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
