package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type TokenType string

const (
	TokenKeywordLet      TokenType = "LET"
	TokenEOF             TokenType = "EOF"
	TokenILLEGAL         TokenType = "ILLEGAL"
	TokenError           TokenType = "ERROR"
	TokenNumber          TokenType = "NUMBER"
	TokenPlus            TokenType = "PLUS"
	TokenMinus           TokenType = "MINUS"
	TokenMultiply        TokenType = "MULTIPLY"
	TokenDivide          TokenType = "DIVIDE"
	TokenLeftParen       TokenType = "LEFT_PAREN"
	TokenRightParen      TokenType = "RIGHT_PAREN"
	TokenIdentifier      TokenType = "IDENTIFIER"
	TokenKeywordIf       TokenType = "IF"
	TokenKeywordElse     TokenType = "ELSE"
	TokenSemiColon       TokenType = "SEMICOLON"
	TokenFunction        TokenType = "FUNCTION"
	TokenComment         TokenType = "COMMENT"
	TokenAssign          TokenType = "ASSIGN"
	TokenLessThan        TokenType = "LESS_THAN"
	TokenGreaterThan     TokenType = "GREATER_THAN"
	TokenColon           TokenType = "COLON"
	TokenComma           TokenType = "COMMA"
	TokenLeftBrace       TokenType = "LEFT_BRACE"
	TokenRightBrace      TokenType = "RIGHT_BRACE"
	TokenReturn          TokenType = "RETURN"
	TokenString          TokenType = "STRING"
	TokenPakage          TokenType = "PACKAGE"
	TokenBuiltInFunction TokenType = "PRINTLINE"
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

// NextToken method reads the input character by character
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.current >= len(l.input) {
		return Token{Type: TokenEOF}
	}
	ch := l.currentChar()

	switch {
	case ch == '=':
		l.advance()
		return Token{Type: TokenAssign, Value: "="}
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
	case ch == '"':
		return l.scanString()
	default:
		fmt.Println(Token{Type: TokenILLEGAL})
		os.Exit(0)
	}

	return Token{Type: TokenError}
}

// currentChar returns the character at the current index
func (l *Lexer) currentChar() rune {
	return rune(l.input[l.current])
}

// advance increments the index of the input
func (l *Lexer) advance() {
	l.current++
}

// skipWhitespace skips over whitespace characters
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

// scanNumber scans number tokens
func (l *Lexer) scanNumber() Token {
	start := l.current
	for l.current < len(l.input) && unicode.IsDigit(l.currentChar()) {
		l.advance()
	}
	return Token{Type: TokenNumber, Value: l.input[start:l.current]}
}

// scanIdentifier scans identifiers and keywords
func (l *Lexer) scanIdentifier() Token {
	start := l.current
	for l.current < len(l.input) && unicode.IsLetter(l.currentChar()) {
		l.advance()
	}
	identifier := l.input[start:l.current]
	if (strings.Contains(l.input, "let") && (strings.Fields(l.input)[0] != "let")) || (strings.Contains(l.input, "Func") && (strings.Fields(l.input)[0] != "Func")) || (strings.Contains(l.input, "func") && (strings.Fields(l.input)[0] != "func")) {
		return Token{Type: TokenError, Value: "ERROR"}
	}

	if identifier == "let" {
		return Token{Type: TokenKeywordLet, Value: identifier}
	} else if identifier == "func" || identifier == "Func" {
		return Token{Type: TokenFunction, Value: "FUNCTION"}
	} else if identifier == "Printline" {
		return Token{Type: TokenBuiltInFunction, Value: identifier}
	}

	return Token{Type: TokenIdentifier, Value: identifier}
}

// scanString scans string literals
func (l *Lexer) scanString() Token {
	if l.currentChar() == '"' {
		l.current++
		start := l.current
		for l.current < len(l.input) && l.input[l.current] != '"' {
			l.current++
		}
		l.current++
		return Token{Type: TokenString, Value: l.input[start : l.current-1]}
	}
	return Token{Type: TokenILLEGAL, Value: ""}
}

func main() {
	if len(os.Args) != 1 {
		fmt.Printf("Usage: go run .\n")
		os.Exit(1)
	}

	file, err := os.Open("../main.ksm")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.Contains(input, "main library") {
			continue
		} else if input == "" {
			continue
		}
		lexer := NewLexer(input)
		fmt.Printf("Statement: %s\n\n", input)
		for {
			token := lexer.NextToken()
			if token.Type == TokenEOF {
				break
			}
			fmt.Printf("Token:%s ,Value:%s\n", token.Type, token.Value)
		}
	}
	fmt.Println(Token{Type: TokenEOF})
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
