# PARSER

## Definition

- A parser is a software component that takes input data (typically text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure, giving a structural representation of the input while checking for correct syntax.

### Responsibilities:

- Reads tokens produced by the lexer.
- Constructs the AST or parse tree based on predefined grammar rules.
- Validates the syntactical correctness of the input.

### Key Methods

``parse()``: Initiates the parsing process.

``expression()``: Parses expressions based on precedence.

``statement()``: Parses different types of statements (e.g., if, loops).

``program()``: Parses the entire program structure.
