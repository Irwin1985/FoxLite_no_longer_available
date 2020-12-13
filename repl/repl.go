// repl/repl.go
// descrip: consola de FoxLite. REPL (Read, Eval, Print, Loop)
// creado: 12/12/2020
// autor: Irwin R.
// colaboradores: <none>
package repl

import (
	"FoxLite/lexer"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		tokens := l.Tokenize()

		for _, tok := range tokens {
			fmt.Printf("%+v\n", tok)
		}
	}
}
