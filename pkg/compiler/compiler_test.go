package compiler

import (
	"testing"
)

func TestCompile(t *testing.T) {
	input := "+-><,."
	expected := []*Instruction{
		&Instruction{INC, 1},
		&Instruction{DEC, 1},
		&Instruction{INCPTR, 1},
		&Instruction{DECPTR, 1},
		&Instruction{GETCHAR, nil},
		&Instruction{PUTCHAR, nil},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileLoops(t *testing.T) {
	input := `+[+[+]+]+`
	expected := []*Instruction{
		&Instruction{INC, 1},
		&Instruction{JZBRACKET, 7},
		&Instruction{INC, 1},
		&Instruction{JZBRACKET, 5},
		&Instruction{INC, 1},
		&Instruction{JNZBRACKET, 3},
		&Instruction{INC, 1},
		&Instruction{JNZBRACKET, 1},
		&Instruction{INC, 1},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileEverything(t *testing.T) {
	input := `+++[---[+]>>>]<<<`
	expected := []*Instruction{
		&Instruction{INC, 1},
		&Instruction{INC, 1},
		&Instruction{INC, 1},
		&Instruction{JZBRACKET, 13},
		&Instruction{DEC, 1},
		&Instruction{DEC, 1},
		&Instruction{DEC, 1},
		&Instruction{JZBRACKET, 9},
		&Instruction{INC, 1},
		&Instruction{JNZBRACKET, 7},
		&Instruction{INCPTR, 1},
		&Instruction{INCPTR, 1},
		&Instruction{INCPTR, 1},
		&Instruction{JNZBRACKET, 3},
		&Instruction{DECPTR, 1},
		&Instruction{DECPTR, 1},
		&Instruction{DECPTR, 1},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}
