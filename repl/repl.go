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
