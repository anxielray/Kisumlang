# _*LEXER*_

- Lexing is the process of breaking down statement blocks of code into  smaller portions called tokens that are able to be unerstood by our interpreter.

## Lexing

- Lexing of code undergoes 2 main processess. First we break down the source code into tokens, a process known as lexical analysis and it is done by a scanner/tokenizer.
- Tokens are small, easily categorizable data structures that are fed to the parser, which does the second transformation turning the tokens into Abstract Syntax Tree.

- Example:
  - If we have a statement such as the ine below and pass it through the lexer;

```sh
"let x = 5 + 5;"
```

expected output:

```sh
[
    LET,
    IDENTIFIER: x,
    ASSIGN,
    INTEGER: 5,
    PLUS_SIGN,
    INTEGER: 5,
    SEMICOLON
]
```

The above example shows the tokens enumerated from the code snippet statement above.

## Defining the tokens

- First things first, we will define the tokens that our lexer will be displaying a s output. Like in the above example.

## Implementing the lexer

- 