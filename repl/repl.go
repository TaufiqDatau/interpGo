package repl

import (
	"bufio"
	"fmt"
	"interpGo/lexer"
	"interpGo/token"
	"io"
)

const PROMPT = ">>"
func Start(in io.Reader, ou io.Writer){
  scanner := bufio.NewScanner(in)

  for{
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()

    if !scanned{
      return
    }

    line := scanner.Text()

    if line == "exit" || line == "quit" {
        fmt.Println("Exiting REPL...")
        return
    }
    l := lexer.New(line)

    for tok:= l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
      fmt.Printf("%+v\n", tok)
    }
  }
}
