package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ..
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LANGLE = "<"
	RANGLE = ">"

	EXCLAMATION = "!"
	SLASH = "/"
	ASTERISK = "*"
	EQ = "=="
	NOT_EQ = "!="


	FUNCTION = "FUNCTION"
	LET = "LET"
	ELSE = "ELSE"
	IF = "IF"
	TRUE = "TRUE"
	FALSE = "FALSE"
	RETURN = "RETURN"

)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"else": ELSE,
	"if": IF,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
