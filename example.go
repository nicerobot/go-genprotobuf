package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type first struct {
	a int32
	b float64
}

func (f first) doFirst() error {
	return fmt.Errorf("")
}

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	if f, err := parser.ParseFile(fset, "example.go", nil, parser.ImportsOnly); err != nil {
		fmt.Println(err)
		return
	} else {
		ast.Print(fset, f)
	}

}
