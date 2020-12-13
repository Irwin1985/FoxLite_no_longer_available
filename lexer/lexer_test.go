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
	input := `var five = 5

var ten = 10

func add(x, y)
  return x + y
endfunc

var result = add(five, ten)
?result

!-/*5
5 < 10 > 5

if (5 < 10)
  return true
else
  return false
endif

10 == 10
10 != 9

"foobar"
"foo bar"

[1, 2]

if 1 ^ 2 > 1
  return true
elif 2 % 2 == 0
  return false
endif

?"Hola Mundo"
"Hola"$"Hola Mundo"
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.NEWLINE, "NEW_LINE"},
		{token.VAR, "var"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.NEWLINE, "NEW_LINE"},
		{token.FUNC, "func"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "NEW_LINE"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.NEWLINE, "NEW_LINE"},
		{token.ENDFUNC, "endfunc"},
		{token.NEWLINE, "NEW_LINE"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "NEW_LINE"},
		{token.QUEST, "?"},
		{token.IDENT, "result"},
		{token.NEWLINE, "NEW_LINE"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.DIV, "/"},
		{token.MUL, "*"},
		{token.INTEGER, "5"},
		{token.NEWLINE, "NEW_LINE"},
		{token.INTEGER, "5"},
		{token.LESS, "<"},
		{token.INTEGER, "10"},
		{token.GREATER, ">"},
		{token.INTEGER, "5"},
		{token.NEWLINE, "NEW_LINE"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INTEGER, "5"},
		{token.LESS, "<"},
		{token.INTEGER, "10"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "NEW_LINE"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.NEWLINE, "NEW_LINE"},
		{token.ELSE, "else"},
		{token.NEWLINE, "NEW_LINE"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.NEWLINE, "NEW_LINE"},
		{token.ENDIF, "endif"},
		{token.NEWLINE, "NEW_LINE"},
		{token.INTEGER, "10"},
		{token.EQUAL, "=="},
		{token.INTEGER, "10"},
		{token.NEWLINE, "NEW_LINE"},
		{token.INTEGER, "10"},
		{token.NOT_EQ, "!="},
		{token.INTEGER, "9"},
		{token.NEWLINE, "NEW_LINE"},
		{token.STRING, "foobar"},
		{token.NEWLINE, "NEW_LINE"},
		{token.STRING, "foo bar"},
		{token.NEWLINE, "NEW_LINE"},
		{token.LBRACKET, "["},
		{token.INTEGER, "1"},
		{token.COMMA, ","},
		{token.INTEGER, "2"},
		{token.RBRACKET, "]"},
		{token.NEWLINE, "NEW_LINE"},
		{token.IF, "if"},
		{token.INTEGER, "1"},
		{token.POW, "^"},
		{token.INTEGER, "2"},
		{token.GREATER, ">"},
		{token.INTEGER, "1"},
		{token.NEWLINE, "NEW_LINE"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.NEWLINE, "NEW_LINE"},
		{token.ELIF, "elif"},
		{token.INTEGER, "2"},
		{token.MOD, "%"},
		{token.INTEGER, "2"},
		{token.EQUAL, "=="},
		{token.INTEGER, "0"},
		{token.NEWLINE, "NEW_LINE"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.NEWLINE, "NEW_LINE"},
		{token.ENDIF, "endif"},
		{token.NEWLINE, "NEW_LINE"},
		{token.QUEST, "?"},
		{token.STRING, "Hola Mundo"},
		{token.NEWLINE, "NEW_LINE"},
		{token.STRING, "Hola"},
		{token.COMP, "$"},
		{token.STRING, "Hola Mundo"},
		{token.NEWLINE, "NEW_LINE"},
		{token.EOF, ""},
	}

	l := New(input)
	tokens := l.Tokenize()

	for i, tt := range tests {
		tok := tokens[i]

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
