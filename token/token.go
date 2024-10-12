package token

/*
Defined the TokenType type to be a string. That allows us to use many different values
as TokenTypes, which in turn allows us to distinguish between different types of tokens.
*/
type TokenType string

type Token struct{
  Type TokenType;
  Literal string
}

const(
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  IDENT   = "IDENT"
  INT     = "INT"
  FLOAT   = "FLOAT"

  //Operators
  ASSIGN  = "="
  PLUS    = "+"
  MINUS   = "-"
  ASTERISK= "*"
  SLASH   = "/"
  BANG    = "!"
  
  //Comparator
  LT      = "<"
  GT      = ">"
  LTE     = "<="
  GTE     = ">="
  EQ      = "=="
  NEQ     = "!="
  
  // Delimiters
  COMMA = ","
  SEMICOLON = ";"
  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  
  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"

)
var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
}

func LookupIndent(indent string) TokenType{
  if tok, ok := keywords[indent]; ok{
    return tok
  }

  return IDENT
}
