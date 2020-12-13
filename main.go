// main.go
// descrip: entrada principal al interprete de FoxLite.
// creado: 12/12/2020
// autor: Irwin R.
// colaboradores: <none>
package main

import (
	"FoxLite/lexer"
	"fmt"
)

func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }
	input := `!5`
	l := lexer.New(input)
	tokens := l.Tokenize()
	for _, token := range tokens {
		fmt.Printf("%+v\n", token)
	}

	// p := parser.New(tokens)
	// program := p.ParseProgram()
	// fmt.Printf("Total stmt: %d", len(program.Statements))

	// if len(program.Statements) != 3 {
	// 	fmt.Println("Error")
	// }
	// fmt.Printf("Hello %s! Wellcome to FoxLite programming language!\n", user.Username)
	// fmt.Printf("Feel free to type some Visual FoxPro commands like\n")
	// repl.Start(os.Stdin, os.Stdout)
}
