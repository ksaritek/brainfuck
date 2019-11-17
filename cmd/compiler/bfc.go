package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ksaritek/brainfuck/pkg/compiler"
)

func main() {
	f := os.Args[1]
	code, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	c := compiler.NewCompiler(string(code))
	instructions := c.Compile()

	bf, err := os.Create(fmt.Sprintf("%s.bytecode", f))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}
	defer bf.Close()

	for _, i := range instructions {
		bf.WriteString(fmt.Sprintf("%s\n", i.String()))
	}
}
