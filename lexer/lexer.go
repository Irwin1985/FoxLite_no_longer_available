// lexer/lexer_test.go
// descrip: convierte el código fuente en un slice de `token`
// creado: 12/12/2020
// autor: Irwin R.
// colaboradores: <none>
package lexer

import (
	"FoxLite/token"
	"strings"
)

type Lexer struct {
	input      string // El código fuente
	curCharPos int
	line       int
	column     int
}

// New Crea una instancia del Lexer
func New(input string) *Lexer {
	return &Lexer{
		input: input + " ",
	}
}

// Tokenize separa el código fuente en un slice de `token`
func (l *Lexer) Tokenize() []token.Token {
	var tokens []token.Token
	state := DEFAULT
	var tokenStr string
	l.line = 1
	l.column = -1
	for index := 0; index < len(l.input); index++ {
		chr := l.input[index]
		l.curCharPos = index
		l.column += 1

		switch state {
		case DEFAULT:
			if isLetter(chr) {
				state = WORD
				tokenStr = string(chr)
			} else if isDigit(chr) {
				state = NUMBER
				tokenStr = string(chr)
			} else if chr == '"' {
				// TODO: incluir soporte para el operador "'"
				state = STRING
			} else if chr == '+' {
				tokens = append(tokens, l.newToken(token.PLUS, "+"))
			} else if chr == '-' {
				tokens = append(tokens, l.newToken(token.MINUS, "-"))
			} else if chr == '*' {
				tokens = append(tokens, l.newToken(token.MUL, "*"))
			} else if chr == '/' {
				tokens = append(tokens, l.newToken(token.DIV, "/"))
			} else if chr == '%' {
				tokens = append(tokens, l.newToken(token.MOD, "%"))
			} else if chr == '^' {
				tokens = append(tokens, l.newToken(token.POW, "^"))
			} else if chr == '$' {
				tokens = append(tokens, l.newToken(token.COMP, "$"))
			} else if chr == '?' {
				tokens = append(tokens, l.newToken(token.QUEST, "?"))
			} else if chr == '(' {
				tokens = append(tokens, l.newToken(token.LPAREN, "("))
			} else if chr == ')' {
				tokens = append(tokens, l.newToken(token.RPAREN, ")"))
			} else if chr == '[' {
				tokens = append(tokens, l.newToken(token.LBRACKET, "["))
			} else if chr == ']' {
				tokens = append(tokens, l.newToken(token.RBRACKET, "]"))
			} else if chr == ',' {
				tokens = append(tokens, l.newToken(token.COMMA, ","))
			} else if chr == ';' {
				tokens = append(tokens, l.newToken(token.SEMICOLON, ";"))
			} else if chr == '<' {
				tok, ok := l.getBinaryOperator(chr, token.LESS, token.LESS_EQ)
				if ok {
					index++
				}
				tokens = append(tokens, tok)
			} else if chr == '>' {
				tok, ok := l.getBinaryOperator(chr, token.GREATER, token.GREATER_EQ)
				if ok {
					index++
				}
				tokens = append(tokens, tok)
			} else if chr == '=' {
				tok, ok := l.getBinaryOperator(chr, token.ASSIGN, token.EQUAL)
				if ok {
					index++
				}
				tokens = append(tokens, tok)
			} else if chr == '!' {
				tok, ok := l.getBinaryOperator(chr, token.BANG, token.NOT_EQ)
				if ok {
					index++
				}
				tokens = append(tokens, tok)
			} else if chr == '\n' {
				// Solo registro el primer salto de línea. El resto es ignorado.
				if len(tokens) > 0 {
					if tokens[len(tokens)-1].Type != token.NEWLINE {
						tokens = append(tokens, l.newToken(token.NEWLINE, "NEW_LINE"))
					}
				}
				l.line += 1
				l.column = 0
			}
		case WORD:
			if !isLetter(chr) {
				// TODO: es una palabra reservada o un identificador?
				tokenType := token.LookupIdent(tokenStr)
				tokens = append(tokens, l.newToken(tokenType, tokenStr))
				state = DEFAULT // De vuelta a DEFAULT.
				tokenStr = ""   // Inicializada para capturar otro Literal.
				index--         // Devolvemos el caracter para no perderlo.
			} else {
				tokenStr += string(chr)
			}
		case NUMBER:
			if isDigit(chr) || chr == '.' {
				tokenStr += string(chr)
			} else {
				var tokenType token.TokenType = token.INTEGER
				// Verificamos el tipo de token (INTEGER ? DOUBLE ?)
				if strings.Contains(tokenStr, ".") {
					tokenType = token.DOUBLE
				}
				tokens = append(tokens, l.newToken(tokenType, tokenStr))
				state = DEFAULT
				tokenStr = ""
				index--
			}
		case STRING:
			if chr == '\\' {
				// TODO: incluir soporte para escape de caracteres.
				// scapeCharacters()
			} else if chr != '"' {
				tokenStr += string(chr)
			} else if chr == '"' {
				// Guardamos el token
				tokens = append(tokens, l.newToken(token.STRING, tokenStr))
				tokenStr = ""
				state = DEFAULT
			}
		default:
			tokens = append(tokens, l.newToken(token.ILLEGAL, string(chr)))
		}
	}
	tokens = append(tokens, l.newToken(token.EOF, ""))
	return tokens
}

// getBinaryOperator analiza los operadores binarios como: '<=', '>=', '==' etc
func (l *Lexer) getBinaryOperator(chr byte, first token.TokenType, second token.TokenType) (token.Token, bool) {
	tokenType := first
	tokenLit := string(chr)
	advance := false
	if l.peek() == '=' {
		tokenType = second
		tokenLit += "="
		advance = true
	}
	return l.newToken(tokenType, tokenLit), advance
}

// peek Muestra el siguiente caracter sin afectar la posición actual.
func (l *Lexer) peek() byte {
	return l.input[l.curCharPos+1]
}

// isLetter verifica si el caracter actual es una letra o un '_'
func isLetter(chr byte) bool {
	return 'a' <= chr && chr <= 'z' || 'A' <= chr && chr <= 'Z' || chr == '_'
}

// isDigit verifica si el caracter actual es un número [0-9]
func isDigit(chr byte) bool {
	return '0' <= chr && chr <= '9'
}

// newToken Crea y devuelve una instancia de `token`
func (l *Lexer) newToken(tok token.TokenType, lit string) token.Token {
	return token.Token{Type: tok, Literal: lit, Line: l.line, Column: l.column}
}

const (
	DEFAULT = iota
	STRING
	NUMBER
	WORD
)
