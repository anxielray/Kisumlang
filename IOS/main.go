package main

import "fmt"

/*
An Internal Object system in interpreters or compilers is a structured way to represent
data values, objects, or even functions during the runtime of a program. The IOS enables
the interpreter to manipulate the objects, perform operations, or manage memory with the
language's execution environment. It acts as a bridge between  the syntax(code) and
the actual computation/ operations.

OBJECTS: Every value like (functions, numbers, strings etc) are objects and each object
has a type (like integer, float, string, etc.)

MEMORY MANAGEMENT: Objects are created, stored in memory, ad accessed through references.

OPERATIONS: Objects support operations like arithmetic operations or string concatenation
handled internally

GARBAGE COLLECTION: Unused objects are cleaned up by garbage collection or explicit memory
management.

Steps:
1. Define the basic basic structures of objects and types.
2. Implement operations that the system can perform on those objects.
3. Create a small evaluator that can perform operations on these objects.
*/

const (
	OBJ_INT   = "INTEGER"
	OBJ_ERROR = "ERROR"
)

type Object interface {
	Type() string
	Inspect() string
}

type Integer struct {
	Value int
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() string {
	return OBJ_INT
}

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (e *Error) Type() string {
	return OBJ_ERROR
}

func Eval(left Object, operator string, right Object) Object {
	if left.Type() == OBJ_INT && right.Type() == OBJ_INT {
		leftVal := left.(*Integer).Value
		rightVal := right.(*Integer).Value
		switch operator {
		case "+":
			return &Integer{Value: leftVal + rightVal}
		case "*":
			return &Integer{Value: leftVal * rightVal}
		default:
			return &Error{Message: "Unknown operator: " + operator}
		}
	}
	return &Error{Message: "type mismatch: " + left.Type() + "and" + right.Type()}
}

func main() {
	left := &Integer{Value: 5}
	right := &Integer{Value: 3}
	result := Eval(left, "+", right)
	fmt.Println("5 + 3 =", result.Inspect())
	result = Eval(left, "*", right)
	fmt.Println("5 * 3 =", result.Inspect())
	result = Eval(left, "-", right)
	fmt.Println(result.Inspect())
}
