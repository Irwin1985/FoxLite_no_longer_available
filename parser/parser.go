// parser/parser.go
// descrip: toma los tokens y produce un árbol de sintaxis abstracta como salida.
// creado: 12/12/2020
// autor: Irwin R.
// colaboradores: <none>
package parser

import (
	"FoxLite/ast"
	"FoxLite/token"
	"fmt"
	"strconv"
)

// Funciones para el analisis de las expresiones normales y binarias
// las expresiones normales son las prefijas y las binarias son las infijas.
type (
	singleExprFn func() ast.Expression
	binaryExprFn func(ast.Expression) ast.Expression
)

// Tipos de operadores relacionados con el análisis de expresiones
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // < ó >
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x ó !x
	CALL        // myFunction(X)
)

// Parser es el encargado de crear los asts.
type Parser struct {
	tokens           []token.Token
	curTokenPosition int      // posición actual del slice de token
	errors           []string // Lista de posibles errores.
	curToken         token.Token
	peekToken        token.Token
	// diccionario de funciones prefijas e infijas.
	singleExprMap map[token.TokenType]singleExprFn
	binaryExprMap map[token.TokenType]binaryExprFn
}

// Registra una función en el diccionario de expresiones simples.
func (p *Parser) registerSingleExpr(tokenType token.TokenType, funName singleExprFn) {
	p.singleExprMap[tokenType] = funName
}

// Registra una función en el diccionario de expresiones binarias.
func (p *Parser) registerBinaryExpr(tokenType token.TokenType, funName binaryExprFn) {
	p.binaryExprMap[tokenType] = funName
}

// New crea una instancia de Parser
func New(t []token.Token) *Parser {
	p := &Parser{tokens: t, errors: []string{}}
	p.curTokenPosition = -1 // el slice de tokens va de 0 a n-1
	p.nextToken()           // Avanzamos al primer token.

	// Registro de funciones de análisis
	p.singleExprMap = make(map[token.TokenType]singleExprFn)
	p.registerSingleExpr(token.IDENT, p.parseIdentifier)
	p.registerSingleExpr(token.INTEGER, p.parseIntegerLiteral)
	p.registerSingleExpr(token.BANG, p.parsePrefixExpression)
	p.registerSingleExpr(token.MINUS, p.parsePrefixExpression)

	return p
}

// nextToken avanza el puntero de los token.
func (p *Parser) nextToken() {
	p.curTokenPosition += 1
	p.curToken = p.getToken(p.curTokenPosition)
	p.peekToken = p.getToken(p.curTokenPosition + 1)
}

// getToken devuelve un token dada su posición (token.EOF si es mayor)
func (p *Parser) getToken(offset int) token.Token {
	if offset >= len(p.tokens) {
		return token.Token{Type: token.EOF, Literal: " "}
	}
	return p.tokens[offset]
}

// Errors devuelve la lista de errores.
func (p *Parser) Errors() []string {
	return p.errors
}

// expectPeek comprueba que los tipos de tokens sean iguales y avanza al siguiente token.
func (p *Parser) expectPeek(ttype token.TokenType) bool {
	if p.peekTokenIs(ttype) {
		p.nextToken()
		return true
	} else {
		p.peekError(ttype)
		return false
	}
}

// peekError mensaje de error solicitado por el método `expectPeek()`
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be '%s', got '%s' instead.",
		t, p.peekToken.Literal)
	p.errors = append(p.errors, msg)
}

// curTokenIs verifica si el token dado es igual al token actual
func (p *Parser) curTokenIs(ttype token.TokenType) bool {
	return p.curToken.Type == ttype
}

// peekTokenIs verifica si el token dado es igual al siguiente token
func (p *Parser) peekTokenIs(ttype token.TokenType) bool {
	return p.peekToken.Type == ttype
}

// ParseProgram crea un ast.Program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement analiza una sentencia y crea su ast.Statement{}
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseVarStatement = `var` identifier `=` expression
func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken} // token `VAR`
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // token `IDENT`

	if !p.expectPeek(token.ASSIGN) { // token `ASSIGN`
		return nil
	}
	// TODO: nos saltamos el analisis de la expresión
	// hasta que encontremos el final de la línea
	for !p.curTokenIs(token.NEWLINE) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement = `return` expression
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken() // Avanzo al siguiente token.

	// TODO: nos saltamos la expresión por el momento.
	for !p.curTokenIs(token.NEWLINE) {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement analiza una expresión en Foxlite
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.curTokenIs(token.NEWLINE) {
		p.nextToken()
	}

	return stmt
}

// parseExpression
func (p *Parser) parseExpression(precedende int) ast.Expression {
	// Buscamos si existe una función asociada al token actual
	// la buscamos como expresión normal o sencilla.
	singleExpr := p.singleExprMap[p.curToken.Type]
	if singleExpr == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExpr := singleExpr()

	return leftExpr
}

// noPrefixParseFnError
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for '%s' found", t)
	p.errors = append(p.errors, msg)
}

// parseIdentifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseIntegerLiteral
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

// parsePrefixExpression
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken() // avanza el operador

	expression.Right = p.parseExpression(PREFIX)

	return expression
}
