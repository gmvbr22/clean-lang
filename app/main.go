package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gmvbr/clean-lang/app/listener"
	"github.com/gmvbr/clean-lang/app/parser"
)


func main() {
	fmt.Printf("Parse: %s\n", os.Args[1])

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewLangLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewLangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	
	antlr.ParseTreeWalkerDefault.Walk(listener.NewTreeLangListener(), p.Program())
}
