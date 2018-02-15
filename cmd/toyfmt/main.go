package main

import (
	"flag"
	"fmt"
	"github.com/go-li/mapast"
	"github.com/go-li/mapast/convert"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

func Printer(s string) {
	if len(s) == 0 {
		fmt.Println()
	} else {
		fmt.Print(s)
	}
}

func main() {
	var filename string
	flag.StringVar(&filename, "I", "", "go source code file to translate")
	flag.Parse()
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "demo", content, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		os.Exit(4)
	}
	asttree := make(map[uint64][]byte)
	ast.Walk(convert.NewConversion(asttree, 0, content), file)
	if false {
		mapast.Dump(Printer, asttree, 0, 0)
		fmt.Println("---------------------------------------------------------")
	}
	mapast.Code(Printer, asttree, 0, 0)
}
