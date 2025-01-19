package token

type TokenType byte

const (
	LET TokenType = iota
	FUNCTION
	IF
	ELSE
	WHILE
	FOR
	BREAK
	CONTINUE
	NUMBER
	STRING
	VARIABLE
	PLUS
	MINUS
	ASTERISK
	SLASH
	MOD
	ASSIGN
	BANG
	EQUALS
	NOT_EQUALS
	LESS_THAN
	GREATER_THAN
	LESS_THAN_OR_EQUAL
	GREATER_THAN_OR_EQUAL
	COMMA
	SEMICOLON
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_CURLY
	RIGHT_CURLY
	LEFT_BRACKET
	RIGHT_BRACKET
	ILLEGAL
	EOF
)

func (t TokenType) String() string {
	return []string{
		"LET",
		"FUNCTION",
		"IF",
		"ELSE",
		"WHILE",
		"FOR",
		"BREAK",
		"CONTINUE",
		"NUMBER",
		"STRING",
		"VARIABLE",
		"PLUS",
		"MINUS",
		"ASTERISK",
		"SLASH",
		"MOD",
		"ASSIGN",
		"BANG",
		"EQUALS",
		"NOT_EQUALS",
		"LESS_THAN",
		"GREATER_THAN",
		"LESS_THAN_OR_EQUAL",
		"GREATER_THAN_OR_EQUAL",
		"COMMA",
		"SEMICOLON",
		"LEFT_PAREN",
		"RIGHT_PAREN",
		"LEFT_CURLY",
		"RIGHT_CURLY",
		"LEFT_BRACKET",
		"RIGHT_BRACKET",
		"ILLEGAL",
		"EOF"}[t]
}

type Token struct {
	Literal  string
	Type     TokenType
	Position int
	Line     int
}

var Keywords = map[string]TokenType{
	"let":      LET,
	"function": FUNCTION,
	"if":       IF,
	"else":     ELSE,
	"while":    WHILE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
}

func LookUpIdent(str string) TokenType {
	if tok, ok := Keywords[str]; ok {
		return tok
	}
	return VARIABLE
}
