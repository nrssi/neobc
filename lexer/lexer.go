package lexer

import "neobc/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.line = 1
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQUALS, Literal: string(ch) + string(l.ch), Position: l.position - 1, Line: l.line}
		} else {
			tok = newToken(token.ASSIGN, string(l.ch), l.position, l.line)
		}
	case '+':
		tok = newToken(token.PLUS, string(l.ch), l.position, l.line)
	case '-':
		tok = newToken(token.MINUS, string(l.ch), l.position, l.line)
	case '!':
		if l.peekChar() == '=' {
			ch := string(l.ch)
			l.readChar()
			tok = newToken(token.NOT_EQUALS, string(ch)+string(l.ch), l.position, l.line)
		} else {
			tok = newToken(token.BANG, string(l.ch), l.position, l.line)
		}
	case '/':
		tok = newToken(token.SLASH, string(l.ch), l.position, l.line)
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch), l.position, l.line)
	case '%':
		tok = newToken(token.MOD, string(l.ch), l.position, l.line)
	case '<':
		if l.peekChar() == '=' {
			ch := string(l.ch)
			l.readChar()
			tok = newToken(token.LESS_THAN_OR_EQUAL, string(ch)+string(l.ch), l.position, l.line)
		} else {
			tok = newToken(token.LESS_THAN, string(l.ch), l.position, l.line)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := string(l.ch)
			l.readChar()
			tok = newToken(token.GREATER_THAN_OR_EQUAL, string(ch)+string(l.ch), l.position, l.line)
		} else {
			tok = newToken(token.GREATER_THAN, string(l.ch), l.position, l.line)
		}
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch), l.position, l.line)
	case ',':
		tok = newToken(token.COMMA, string(l.ch), l.position, l.line)
	case '{':
		tok = newToken(token.LEFT_CURLY, string(l.ch), l.position, l.line)
	case '}':
		tok = newToken(token.RIGHT_CURLY, string(l.ch), l.position, l.line)
	case '(':
		tok = newToken(token.LEFT_PAREN, string(l.ch), l.position, l.line)
	case ')':
		tok = newToken(token.RIGHT_PAREN, string(l.ch), l.position, l.line)
	case '[':
		tok = newToken(token.LEFT_BRACKET, string(l.ch), l.position, l.line)
	case ']':
		tok = newToken(token.RIGHT_BRACKET, string(l.ch), l.position, l.line)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			tok.Position = l.position - len(tok.Literal)
			tok.Line = l.line

			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumber()
			tok.Position = l.position - len(tok.Literal)
			tok.Line = l.line
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch), l.position, l.line)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch string, pos, line int) token.Token {
	return token.Token{Type: tokenType, Literal: ch, Position: pos, Line: line}
}
