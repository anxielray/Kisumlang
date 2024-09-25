package main

/*
An AST is a hierarchal tree structure that represents the structure of the source code.
Each node in the tree represents a language construct, such as an operation, expression, or statement.
Unlike a parser tree (which directly mirrors the grammar of the language), an AST abstracts away
unnecessary details like parenthesis and focuses on  the logical construction of the code.

NODES : Each node in the AST represents a  construct lik a  variable, operator, or function call.
INTERNAL NODES: Represents operations and their children representing the operands.
LEAF NODES: Represents the basic elements like numbers, variables, or literals
PURPOSE: ASTs simplify later stages of interpretation or compilation by abstracting away syntactic details
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	TOKEN_INT     = "INT"
	TOKEN_PLUS    = "PLUS"
	TOKEN_MINUS   = "MINUS"
	TOKEN_COMMENT = "COMMENT"
	TOKEN_DIV     = "DIV"
	TOKEN_ILLEGAL = "ILLEGAL"
	TOKEN_MUL     = "MUL"
	TOKEN_EOF     = "EOF"
)

type Token struct {
	typ   string
	value string
}

type Lexer struct {
	text        string
	pos         int
	currentChar byte
}

func NewLexer(text string) *Lexer {
	lexer := &Lexer{text: text, pos: 0}
	lexer.currentChar = text[lexer.pos]
	return lexer
}

func (l *Lexer) Advanced() {
	l.pos++
	if l.pos >= len(l.text) {
		l.currentChar = 0
	} else {
		l.currentChar = l.text[l.pos]
	}
}

func (l *Lexer) Integer() string {
	result := ""
	for l.currentChar != 0 && unicode.IsDigit(rune(l.currentChar)) {
		result += string(l.currentChar)
		l.Advanced()
	}
	return result
}

func (l *Lexer) GetNextToken() Token {
	for l.currentChar != 0 {
		if l.currentChar == ' ' {
			l.SkipWhitespace()
			continue
		}
		if unicode.IsDigit(rune(l.currentChar)) {
			return Token{TOKEN_INT, l.Integer()}
		}
		if l.currentChar == '+' {
			l.Advanced()
			return Token{TOKEN_PLUS, "+"}
		}
		if l.currentChar == '-' {
			l.Advanced()
			return Token{TOKEN_MINUS, "-"}
		}
		if l.currentChar == '*' {
			l.Advanced()
			return Token{TOKEN_MUL, "*"}
		}
		if l.currentChar == '/' {
			l.scanComment()
		}
		fmt.Println("Invalid character")
		os.Exit(1)
	}
	return Token{TOKEN_EOF, ""}
}

type BinOPNode struct {
	left     ASTNode
	operator Token
	right    ASTNode
}

type NumNode struct {
	value int
}

func (b *BinOPNode) Interpret() int {
	switch b.operator.typ {
	case TOKEN_PLUS:
		return b.left.Interpret() + b.right.Interpret()
	case TOKEN_MUL:
		return b.left.Interpret() * b.right.Interpret()
	default:
		fmt.Println("Unknown operator")
		os.Exit(1)
	}
	return 0
}

func (n *NumNode) Interpret() int {
	return n.value
}

func (p *Parser) Factor() ASTNode {
	value, _ := strconv.Atoi(p.token.value)
	node := &NumNode{value: value}
	p.Eat(TOKEN_INT)
	return node
}

type ASTNode interface {
	Interpret() int
}

func (p *Parser) Term() ASTNode {
	node := p.Factor()

	for p.token.typ == TOKEN_MUL {
		operator := p.token
		p.Eat(TOKEN_MUL)
		node = &BinOPNode{left: node, operator: operator, right: p.Factor()}
	}
	return node
}

func (p *Parser) AExpr() ASTNode {
	node := p.Term()
	for p.token.typ == TOKEN_PLUS {
		operator := p.token
		p.Eat(TOKEN_PLUS)
		node = &BinOPNode{left: node, operator: operator, right: p.Term()}
	}
	return node
}

func (l *Lexer) SkipWhitespace() {
	for l.currentChar != 0 && l.currentChar == ' ' {
		l.Advanced()
	}
}

type Parser struct {
	lexer *Lexer
	token Token
}

func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.token = lexer.GetNextToken()
	return parser
}

func (p *Parser) Eat(tokenType string) {
	if p.token.typ == tokenType {
		p.token = p.lexer.GetNextToken()
	} else {
		fmt.Println("Syntax error")
	}
}

func (l *Lexer) scanComment() Token {
	if l.pos < len(l.text) {
		var count int
		for _, c := range l.text {
			if c == '/' {
				count++
			} else {
				continue
			}
		}
		// println(count)
		if count == 1 {
			l.pos++
			return Token{typ: TOKEN_DIV, value: "/"}
		} else if count > 1 {
			start := l.pos
			for l.pos < len(l.text) && l.text[l.pos] != '\n' {
				l.pos++
			}
			return Token{typ: TOKEN_COMMENT, value: l.text[start:l.pos]}
		}
	}
	// If no comment or divide token is found, return an illegal token
	return Token{typ: TOKEN_ILLEGAL}
}

func (p *Parser) Expr() int {
	left, _ := strconv.Atoi(p.token.value)
	p.Eat(TOKEN_INT)

	if p.token.typ == TOKEN_PLUS {
		p.Eat(TOKEN_PLUS)
		right, _ := strconv.Atoi(p.token.value)
		p.Eat(TOKEN_INT)
		return left + right
	} else if p.token.typ == TOKEN_MINUS {
		p.Eat(TOKEN_MINUS)
		right, _ := strconv.Atoi(p.token.value)
		p.Eat(TOKEN_INT)
		return left - right
	}
	fmt.Println("Invalid syntax")
	return 0
}

func main() {
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
		parser := NewParser(lexer)
		ast := parser.AExpr()
		result := ast.Interpret()
		fmt.Println("Result: ", result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
