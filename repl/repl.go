package repl

import (
	"bufio"
	"fmt"
	"io"

	"interpreter-go.com/lexer"
	"interpreter-go.com/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
			fmt.Printf("%+v\n", tok)
		}
	}
}
