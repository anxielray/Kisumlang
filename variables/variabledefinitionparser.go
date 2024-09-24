package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	// "kisumulang/lexer"
)



type VariableTokens struct {
	TokenKeywordLet string
	VarName    string
	VarType    string // Should be of variable type
	AssignmentOperator string
	Value string
	Line int
}
// create an instance of the variable tokens

func NewToken(tokenKeywordLet, varName, VarType, assignmentOperator, value string, line int) VariableTokens{
	return VariableTokens{
		TokenKeywordLet : tokenKeywordLet,
		VarName: varName,
		VarType: VarType,
		AssignmentOperator: assignmentOperator,
		Value: value,
		Line: line,
	}
}

//Consumes characters which form lexeme and gives back a token
func Scanner(line string) []string{
	pattern := regexp.MustCompile(`^let\s([a-zA-z_][a-zA-Z0-9_]*)\s([a-zA-Z][a-zA-Z0-9_]*)\s*(=)\s*(.+)$`)
	if !pattern.MatchString(line){
		fmt.Println("Error reading the line")
		
	}
	tokenMap := make(map[string]string)
	if matches := pattern.FindStringSubmatch(line); matches != nil{
		tokenMap["keyword"] = "let"
		tokenMap["varName"] = matches[1]
		tokenMap["varType"] = matches[2]
		tokenMap["assignmentOperator"] = matches[3]
		tokenMap["value"] = matches[4]
	}

	word  := ""
	var tokenSlice []string

	for i, char := range line{
		if char != ' '{
			word += string(char)
		}else {
			if word != ""{
				tokenSlice = append(tokenSlice, word)
				word = ""
			}
		}
		if i == len(line)-1 && word != ""{
			tokenSlice = append(tokenSlice, word)
		}
	}

	return tokenSlice
}


// Declaring the error form in variable declarations
func variableErrors(expected, got string) error {
	fmt.Printf("Expecting %v, got %v\n", expected, got)
	return errors.New("Syntax error")
}




func main() {
	file, err := os.Open("test.ksm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens:= Scanner(line)
		
		fmt.Println(tokens)
	}
	if err := scanner.Err(); err != nil{
		fmt.Println(err)
	}
}
