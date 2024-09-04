# _*Pseudo-code*_

- This pseudo-code outlines a basic structure and functionality for the Kisum programming language with a focus on essential data structures and their methods
- This document outlines the key steps and considerations for developing an interpreted programming language.

  ## _*Overview*_

- Develop an interpreted programming language called Kisumu. Kisumu will have a syntax similar to Go, support basic data structures, and feature a ".ksm" file extension.

## _*Project Structure*_

```
project/
│
├── main.go           # Entry point for the interpreter
├── lexer.go          # Lexer for tokenizing input
├── parser.go         # Parser for generating the abstract syntax tree (AST)
├── interpreter.go     # Interpreter for executing the AST
├── types.go          # Definition of data structures and methods
├── builtin_functions.go # Built-in functions for standard operations
└── README.md         # Project documentation
```

- ## Data Structures

### 1. Number
- Represents both integer and float values.

#### Methods:
- `Add(a Number): Number` — Returns the sum of two numbers.
- `Subtract(a Number): Number` — Returns the difference of two numbers.
- `Multiply(a Number): Number` — Returns the product of two numbers.

### 2. String
- Represents a sequence of characters.

#### Methods:
- `Length(): int` — Returns the number of characters in the string.
- `Substring(start int, length int): String` — Returns a substring.
- `Concat(other String): String` — Concatenates two strings and returns a new string.

### 3. Boolean
- Represents a truth value (`true` or `false`).

#### Methods:
- `Not(): Boolean` — Returns the negation of the Boolean value.
- `And(other Boolean): Boolean` — Returns the logical AND of two Boolean values.
- `Or(other Boolean): Boolean` — Returns the logical OR of two Boolean values.

### 4. Null
- Represents the absence of value.

#### Methods:
- `IsNull(): Boolean` — Returns `true` if the value is null.
- `ToString(): String` — Returns the string "null".
- `Equals(other Null): Boolean` — Returns `true` if both are null.

### 5. Array
- Represents a collection of elements.

#### Methods:
- `Length(): int` — Returns the size of the array.
- `First(): Element` — Returns the first element of the array.
- `Get(index int): Element` — Returns the element at specified index.

### 6. Object/Hash
- Represents a collection of key-value pairs.

#### Methods:
- `Get(key String): Value` — Returns the value associated with the key.
- `Set(key String, value Value): Null` — Sets a key-value pair in the object.
- `Keys(): Array` — Returns an array of keys in the object.

## Lexer
- Tokenizes the input source into a series of tokens.
- Example tokens: `NUMBER`, `STRING`, `BOOLEAN`, `NULL`, `ARRAY`, `OBJECT`, `IDENTIFIER`, `PLUS`, `MINUS`, `EQUALS`, etc.

## Parser
- Converts the list of tokens into an Abstract Syntax Tree (AST).
- Supports defining variables, functions, control flow (if statements, loops), and more.

## Interpreter
- Walks through the AST and executes the defined operations in Kisumlang.
- Manages data structures and their methods.

## Built-in Functions
- Provide standard functionalities such as mathematical operations, string manipulation, etc.
- Example: `Print(value Value): Null` — prints the value to the console.

## _*Example Code in Kisumlang*_

```ksm
// Example Kisumu code
var myArray = [1, 2, 3, 4.5]
var size = myArray.Length()  // Calls length method
var firstElement = myArray.First()  // Calls first method
var myString = "Hello"
var stringSize = myString.Length()  // Calls length method
```

## _*Future Enhancements*_

-  Implement error handling and exception management.
-  Add support for more complex data structures (e.g., sets, tuples).
- Improve performance optimizations for the interpreter.
  
