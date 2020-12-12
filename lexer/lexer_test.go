// lexer/lexer_test.go
// descrip: pruebas unitarias para el `lexer`.
// creado: 12/12/2020
// autor: Irwin R.
// colaboradores: <none>
package lexer

import (
	"FoxLite/token"
	"testing"
)

func TestTokenList(t *testing.T) {
	tests := []struct {
		input           string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"+", token.PLUS, "+"},
		{"-", token.MINUS, "-"},
		{"*", token.MUL, "*"},
		{"/", token.DIV, "/"},
		{"^", token.POW, "^"},
		{"$", token.COMP, "$"},
		{"?", token.COMP, "?"},
		{"%", token.MOD, "%"},
		{"<", token.LESS, "<"},
		{">", token.GREATER, ">"},
		{"<=", token.LESS_EQ, "<="},
		{">=", token.GREATER_EQ, ">="},
		{"=", token.ASSIGN, "="},
		{"==", token.EQUAL, "=="},
		{"!=", token.NOT_EQ, "!="},
		{"!", token.BANG, "!"},
		{"(", token.LPAREN, "("},
		{")", token.RPAREN, ")"},
		{"[", token.LBRACKET, "["},
		{"]", token.RBRACKET, "]"},
		{",", token.COMMA, ","},
		{";", token.SEMICOLON, ";"},
		{"and", token.AND, "and"},
		{"or", token.OR, "or"},
		{"not", token.NOT, "not"},
		{"xor", token.XOR, "xor"},
		{"var", token.VAR, "var"},
		{"func", token.FUNC, "func"},
		{"endfunc", token.ENDFUNC, "endfunc"},
		{"five", token.IDENT, "five"},
		{"1985", token.INTEGER, "1985"},
		{"11.35", token.DOUBLE, "11.35"},
		{`"string"`, token.STRING, "string"},
	}

	for _, tt := range tests {
		l := New(tt.input)
		tokens := l.Tokenize()
		if len(tokens) != 1 {
			t.Fatalf("wrong token size. expected=%d, got=%d", 1, len(tokens))
		}
		if tokens[0].Literal != tt.expectedLiteral {
			t.Fatalf("tokens[0].Literal is different. expected=%s, got=%s", tt.expectedLiteral, tokens[0].Literal)
		}
	}
}
