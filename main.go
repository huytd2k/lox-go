package main

import (
	"fmt"
	"os"
)

var lox Lox

func main() {
	lox.start()
}

type Lox struct {
	hadError bool
}

func (lox *Lox) runFromSource(source []byte) {
	if lox.hadError {
		os.Exit(65)
	}
}

func (lox *Lox) runRepl() {
	var input string
	for {
		fmt.Print("> ")
		fmt.Scanln(&input)
		lox.hadError = false
	}
}

func (*Lox) exit(msg string, code int) {
	fmt.Print(msg)
	os.Exit(code)
}

func (lox *Lox) start() {
	args := os.Args[1:]
	argslen := len(args)

	if argslen > 1 {
		lox.exit("Usage lox <script>", 1)
	} else if argslen == 1 {
		source, err := os.ReadFile(args[1])
		if err != nil {
			lox.exit("Error while reading file %s", 1)
		}
		lox.runFromSource(source)
	} else {
		lox.runRepl()
	}
}
