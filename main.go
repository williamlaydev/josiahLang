package main

import (
	"bufio"
	"flag"
	"fmt"
	"josiahLang/evaluator"
	"josiahLang/lexer"
	"josiahLang/object"
	"josiahLang/parser"
	"josiahLang/repl"
	"os"
	"regexp"
)

func main() {
	sourceName := flag.String("s", "", "JosiahLang source file name")
	flag.Parse()

	// Read the file
	if !isValidSourceFile(sourceName) {
		fmt.Printf("Source wrong file type, want=*.josiah got=%s\n", *sourceName)
		return
	}

	file, err := os.Open(*sourceName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sourceCode, err := readFile(file)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// Source code to Go
	env := object.NewEnvironment()
	l := lexer.New(sourceCode)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stdout, p.Errors())
		return
	}

	res := evaluator.Eval(program, env)

	// DEBUGGING TOOL
	if res != nil {
		// io.WriteString(os.Stdout, res.Inspect())
		// io.WriteString(os.Stdout, "\n")
	}
}

func isValidSourceFile(fileName *string) bool {
	match, _ := regexp.MatchString(`\.josiah$`, *fileName)
	return match
}

func readFile(file *os.File) (string, error) {
	var sourceCode string
	scanner := bufio.NewScanner(file)

	// Scan the entire file in
	for scanner.Scan() {
		sourceCode += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return sourceCode, nil
}
