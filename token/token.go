// token/token.go
// descrip: Definición del tipo de dato `token`
// creado: 12/12/2020
// autor: Irwin R.
// colaboradores: <none>

package token

type TokenType int // Tipo de token

// Listado de palabras reservadas.
var keywords = map[string]TokenType{
	"var":     VAR,
	"func":    FUNC,
	"endfunc": ENDFUNC,
	"not":     NOT,
	"and":     AND,
	"or":      OR,
	"xor":     XOR,
}

const (
	EOF TokenType = iota
	UNKNOWN

	// Identificadores + literales
	IDENT
	INTEGER
	DOUBLE
	STRING

	// Operadores Numéricos
	PLUS  // "+"
	MINUS // "-"
	MUL   // "*"
	DIV   // "/"
	POW   // "^"
	MOD   // "%"

	COMP  // "$" comparación de cadenas, ej: "Fox" $ "FoxLite"
	QUEST // "?"

	// Operadores lógicos
	BANG // "NOT" ó "!"
	AND
	OR
	XOR // Literal "xor". Ya no tendrá soporte "#"
	NOT // Literal "not"

	// Operadores relacionales
	LESS    // "<"
	GREATER // ">"
	ASSIGN  // "="
	EQUAL   // "==" EQUAL y ASSIGN son distintos.

	LESS_EQ    // "<="
	GREATER_EQ // ">="
	NOT_EQ     // "!="

	// Delimitadores
	COMMA     // ","
	SEMICOLON // ";"

	// Caracteres especiales y de puntuación
	LPAREN   // "("
	RPAREN   // ")"
	LBRACKET // "["
	RBRACKET // "]"

	// Palabras reservadas (solo funciones por ahora)
	VAR     // "var" será la encargada de enlazar variables. No habrá LOCAL, PRIVATE, PUBLIC
	FUNC    // "func" será la encargada de crear funciones literales. No habrá PROCEDURE
	ENDFUNC // "endfuc" terminador para las funciones literales.
)

type Token struct {
	Type    TokenType // Todo token tiene un tipo. Ej: STRING, INTEGER, BOOL, ETC.
	Literal string    // y un valor asociado. Ej: [STRING, "Hola"], [INTEGER, 10]
	Line    uint16    // La línea donde comienza el token.
	Column  uint16    // La columna donde comienza el token.
}

// LookupIdent busca si el literal es una palabra reservada o un identificador.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
