package vm

import (
	"bytes"
	"github.com/ksaritek/brainfuck/pkg/compiler"
	"testing"
)

func TestIncrement(t *testing.T) {
	c := compiler.NewCompiler("+++++")
	instructions := c.Compile()

	m := NewMachine(instructions, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	if m.memory[0] != 5 {
		t.Errorf("cell not correctly incremented. got=%d", m.memory[0])
	}
}

func TestDecrement(t *testing.T) {
	i := "++++++++++-----"
	c := compiler.NewCompiler(i)
	ins := c.Compile()

	m := NewMachine(ins, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	if m.memory[0] != 5 {
		t.Errorf("cell not correctly decremented. got=%d", m.memory[0])
	}
}

func TestIncrementingDataPointer(t *testing.T) {
	c := compiler.NewCompiler("+>++>+++")
	ins := c.Compile()

	m := NewMachine(ins, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	for i, expected := range []uint8{1, 2, 3} {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestDecrementDataPointer(t *testing.T) {
	c := compiler.NewCompiler(">>+++<++<+")
	ins := c.Compile()

	m := NewMachine(ins, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	for i, expected := range []uint8{1, 2, 3} {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestReadChar(t *testing.T) {
	in := bytes.NewBufferString("ABCDEF")
	out := new(bytes.Buffer)

	c := compiler.NewCompiler(",>,>,>,>,>,>")
	ins := c.Compile()

	m := NewMachine(ins, in, out)

	m.Execute()

	expectedMemory := []uint8{
		uint8('A'),
		uint8('B'),
		uint8('C'),
		uint8('D'),
		uint8('E'),
		uint8('F'),
	}

	for i, expected := range expectedMemory {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestPutChar(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	c := compiler.NewCompiler(".>.>.>.>.>.>")
	ins := c.Compile()

	m := NewMachine(ins, in, out)

	setupMemory := []uint8{
		uint8('A'),
		uint8('B'),
		uint8('C'),
		uint8('D'),
		uint8('E'),
		uint8('F'),
	}

	for i, value := range setupMemory {
		m.memory[i] = value
	}

	m.Execute()

	output := out.String()
	if output != "ABCDEF" {
		t.Errorf("output wrong. got=%q", output)
	}

}

func TestHelloWorld(t *testing.T) {
	type TestData struct {
		s string
		c string
	}

	testData := []TestData{
		{"Hello World!\n", "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."},
	}

	for _, td := range testData {
		in := bytes.NewBufferString("")
		out := new(bytes.Buffer)

		c := compiler.NewCompiler(td.c)
		ins := c.Compile()

		m := NewMachine(ins, in, out)

		m.Execute()

		output := out.String()
		if output != td.s {
			t.Errorf("output wrong. got=%q", output)
		}
	}
}
