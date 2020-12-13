// token/token.go
// descrip: Definición del tipo de dato `token`
// creado: 12/12/2020
// autor: Irwin R.
// colaboradores: <none>

package token

type TokenType string // Tipo de token

// Listado de palabras reservadas.
var keywords = map[string]TokenType{
	"var":     VAR,
	"func":    FUNC,
	"endfunc": ENDFUNC,
	"not":     NOT,
	"and":     AND,
	"or":      OR,
	"xor":     XOR,
	"return":  RETURN,
	"if":      IF,
	"else":    ELSE,
	"elif":    ELIF,
	"endif":   ENDIF,
	"true":    TRUE,
	"false":   FALSE,
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
	NEWLINE = "NEWLINE"

	// Identificadores + literales
	IDENT   = "IDENT"
	INTEGER = "INTEGER"
	DOUBLE  = "DOUBLE"
	STRING  = "STRING"

	// Operadores Numéricos
	PLUS  = "+"
	MINUS = "-"
	MUL   = "*"
	DIV   = "/"
	POW   = "^"
	MOD   = "%"

	COMP  = "$" // comparación de cadenas, ej: "Fox" $ "FoxLite"
	QUEST = "?"

	// Operadores lógicos
	BANG  = "!"
	AND   = "AND"
	OR    = "OR"
	XOR   = "XOR" // Literal "xor". Ya no tendrá soporte "#"
	NOT   = "NOT" // Literal "not"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Operadores relacionales
	LESS    = "<"
	GREATER = ">"
	ASSIGN  = "="
	EQUAL   = "==" //EQUAL y ASSIGN son distintos.

	LESS_EQ    = "<="
	GREATER_EQ = ">="
	NOT_EQ     = "!="

	// Delimitadores
	COMMA     = ","
	SEMICOLON = ";"

	// Caracteres especiales y de puntuación
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Palabras reservadas (solo funciones por ahora)
	VAR     = "VAR"     // será la encargada de enlazar variables. No habrá LOCAL, PRIVATE, PUBLIC
	FUNC    = "FUNC"    // será la encargada de crear funciones literales. No habrá PROCEDURE
	ENDFUNC = "ENDFUNC" // terminador para las funciones literales.
	RETURN  = "RETURN"

	// Condicional
	IF    = "IF"
	ELSE  = "ELSE"
	ELIF  = "ELIF"
	ENDIF = "ENDIF"
)

type Token struct {
	Type    TokenType // Todo token tiene un tipo. Ej: STRING, INTEGER, BOOL, ETC.
	Literal string    // y un valor asociado. Ej: [STRING, "Hola"], [INTEGER, 10]
	Line    int       // La línea donde comienza el token.
	Column  int       // La columna donde comienza el token.
}

// LookupIdent busca si el literal es una palabra reservada o un identificador.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
