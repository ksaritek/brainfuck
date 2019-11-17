package compiler

type Compiler struct {
	code         string
	instructions []*Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		instructions: make([]*Instruction, len(code)),
	}
}

func (c *Compiler) Compile() []*Instruction {
	jmpStack := []int{}

	for i, statement := range c.code {
		switch statement {
		case '[':
			c.instructions[i] = &Instruction{Operator: JZBRACKET, Operand: 0}
			jmpStack = append(jmpStack, i)
		case ']':
			// Pop position of last JZBRACKET ("[") instruction off stack
			openPos := jmpStack[len(jmpStack)-1]
			jmpStack = jmpStack[:len(jmpStack)-1]
			// Emit the new JNZBRACKET ("]") instruction, with correct position as argument
			c.instructions[i] = &Instruction{Operator: JNZBRACKET, Operand: openPos}
			// Patch the old JNZBRACKET ("[") instruction with new position
			c.instructions[openPos].Operand = i

		case '+':
			c.instructions[i] = &Instruction{Operator: INC, Operand: 1}
		case '-':
			c.instructions[i] = &Instruction{Operator: DEC, Operand: 1}
		case '<':
			c.instructions[i] = &Instruction{Operator: DECPTR, Operand: 1}
		case '>':
			c.instructions[i] = &Instruction{Operator: INCPTR, Operand: 1}
		case '.':
			c.instructions[i] = &Instruction{Operator: PUTCHAR, Operand: nil}
		case ',':
			c.instructions[i] = &Instruction{Operator: GETCHAR, Operand: nil}
		default:
			panic("undefined instruction")
		}
	}

	return c.instructions
}
