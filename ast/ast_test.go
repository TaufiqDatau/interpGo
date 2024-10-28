package ast

import (
	"interpGo/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{ 
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "paring"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "paring myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
