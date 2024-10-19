package parser

import (
	"interpGo/ast"
	"interpGo/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
  paring x = 5;
  paring y = 10
  paring endo = 399;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"endo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "paring" {

	}
}
