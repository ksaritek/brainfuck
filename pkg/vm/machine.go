package vm

import (
	"io"

	"github.com/ksaritek/brainfuck/pkg/compiler"
)

type Machine struct {
	code []*compiler.Instruction
	ip   int

	memory [30000]uint8
	dp     int

	input  io.Reader
	output io.Writer

	readBuf []byte
}

func NewMachine(instructions []*compiler.Instruction, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:    instructions,
		input:   in,
		output:  out,
		readBuf: make([]byte, 1),
	}
}

func (m *Machine) Execute() {
	for m.ip < len(m.code) {
		ins := m.code[m.ip]

		switch ins.Operator {
		case compiler.INC:
			m.memory[m.dp] += ins.Operand.(uint8)
		case compiler.DEC:
			m.memory[m.dp] -= ins.Operand.(uint8)
		case compiler.INCPTR:
			m.dp += ins.Operand.(int)
		case compiler.DECPTR:
			m.dp -= ins.Operand.(int)
		case compiler.PUTCHAR:
			m.putChar()
		case compiler.GETCHAR:
			m.getChar()
		case compiler.JZBRACKET:
			if m.memory[m.dp] == 0 {
				m.ip = ins.Operand.(int)
				continue
			}
		case compiler.JNZBRACKET:
			if m.memory[m.dp] != 0 {
				m.ip = ins.Operand.(int)
				continue
			}
		}

		m.ip++
	}
}

func (m *Machine) getChar() {
	n, err := m.input.Read(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes read")
	}

	m.memory[m.dp] = uint8(m.readBuf[0])
}

func (m *Machine) putChar() {
	m.readBuf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes written")
	}
}
