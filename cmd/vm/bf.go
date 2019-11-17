package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/ksaritek/brainfuck/pkg/compiler"
	"github.com/ksaritek/brainfuck/pkg/vm"
)

func main() {
	f := os.Args[1]

	bc, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	is := compiler.LoadInstructions(bytes.NewReader(bc))
	os.Stdout.WriteString(">>> Welcome to BF Interpreter!!! \n")
	m := vm.NewMachine(is, os.Stdin, os.Stdout)

	m.Execute()

	os.Stdout.WriteString("\n")
}
