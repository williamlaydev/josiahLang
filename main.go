package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"josiahLang/evaluator"
	"josiahLang/lexer"
	"josiahLang/object"
	"josiahLang/parser"
	"josiahLang/repl"
	"os"
)

func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Hello %s! This is the REPL for JosiahLang\n",
	// 	user.Username)

	// repl.Start(os.Stdin, os.Stdout)

	sourceName := flag.String("s", "", "JosiahLang source file name")
	flag.Parse()

	// Read the file
	file, err := os.Open(*sourceName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	env := object.NewEnvironment()

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			repl.PrintParserErrors(os.Stdout, p.Errors())
			continue
		}

		res := evaluator.Eval(program, env)
		if res != nil {
			io.WriteString(os.Stdout, res.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}
}
