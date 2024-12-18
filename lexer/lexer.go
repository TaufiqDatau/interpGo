package lexer

import (
	"interpGo/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readCharacter()
	return l
}

func (l *Lexer) readCharacter() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.whiteSpacesIgnorer()

	switch l.ch {
	case '=':
		if l.peekNextCharacter() == '=' {
			l.readCharacter()
			tok.Literal = "=="
			tok.Type = token.EQ
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		if l.peekNextCharacter() == '=' {
			l.readCharacter()
			tok.Literal = "!="
			tok.Type = token.NEQ
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		if l.peekNextCharacter() == '=' {
			l.readCharacter()
			tok.Type = token.LTE
			tok.Literal = "<="
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekNextCharacter() == '=' {
			l.readCharacter()
			tok.Type = token.GTE
			tok.Literal = ">="
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = readIdentifier(l)
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal, tok.Type = readNumber(l)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readCharacter()
	return tok
}

func readIdentifier(l *Lexer) string {
	position := l.position
	for isLetter(l.ch) {
		l.readCharacter()
	}
	return l.input[position:l.position]
}

func readNumber(l *Lexer) (string, token.TokenType) {
	isFloat := false
	position := l.position
	for isDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			isFloat = true
		}
		l.readCharacter()
	}
	var tokenType string
	if isFloat {
		tokenType = token.FLOAT
	} else {
		tokenType = token.INT
	}
	return l.input[position:l.position], token.TokenType(tokenType)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) whiteSpacesIgnorer() {
	for l.ch == ' ' ||
		l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readCharacter()
	}
}

func (l *Lexer) peekNextCharacter() byte {
	nextPos := l.position + 1
	if nextPos > len(l.input) {
		return 0
	}

	return l.input[nextPos]
}
