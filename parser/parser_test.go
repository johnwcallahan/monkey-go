package parser

import (
	"testing"

	"github.com/johnwcallahan/monkey-go/ast"
	"github.com/johnwcallahan/monkey-go/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. Got=%T", s)
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. Got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.Token.Literal != name {
		t.Errorf("s.Name not '%s'. Got=%s", name, letStmt.Name)
	}

	return true
}