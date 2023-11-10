package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"
	TRUE  = "true"
	FALSE = "false"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
