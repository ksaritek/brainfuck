package compiler

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Mnemonic byte

const (
	INC        Mnemonic = '+'
	DEC        Mnemonic = '-'
	INCPTR     Mnemonic = '>'
	DECPTR     Mnemonic = '<'
	PUTCHAR    Mnemonic = '.'
	GETCHAR    Mnemonic = ','
	JZBRACKET  Mnemonic = '['
	JNZBRACKET Mnemonic = ']'
)

type Instruction struct {
	Operator Mnemonic
	Operand  interface{}
}

func LoadInstructions(r io.Reader) []*Instruction {
	scanner := bufio.NewScanner(r)

	instructions := []*Instruction{}

	for scanner.Scan() {
		i := strings.Fields(scanner.Text())
		switch i[0] {
		case "INC":
			if o, err := strconv.Atoi(i[1]); err != nil {
				panic(fmt.Errorf("could not parse INC - %v", i))
			} else {
				instructions = pushInstruction(instructions, INC, uint8(o))
			}
		case "DEC":
			if o, err := strconv.Atoi(i[1]); err != nil {
				panic(fmt.Errorf("could not parse DEC - %v", i))
			} else {
				instructions = pushInstruction(instructions, DEC, uint8(o))
			}
		case "INCPTR":
			if o, err := strconv.Atoi(i[1]); err != nil {
				panic(fmt.Errorf("could not parse INCPTR - %v", i))
			} else {
				instructions = pushInstruction(instructions, INCPTR, o)
			}
		case "DECPTR":
			if o, err := strconv.Atoi(i[1]); err != nil {
				panic(fmt.Errorf("could not parse DECPTR - %v", i))
			} else {
				instructions = pushInstruction(instructions, DECPTR, o)
			}
		case "PUTCHAR":
			instructions = pushInstruction(instructions, PUTCHAR, nil)
		case "GETCHAR":
			instructions = pushInstruction(instructions, GETCHAR, nil)
		case "JZBRACKET":
			if o, err := strconv.Atoi(i[1]); err != nil {
				panic(fmt.Errorf("could not parse JZBRACKET - %v", i))
			} else {
				instructions = pushInstruction(instructions, JZBRACKET, o)
			}
		case "JNZBRACKET":
			if o, err := strconv.Atoi(i[1]); err != nil {
				panic(fmt.Errorf("could not parse JNZBRACKET - %v", i))
			} else {
				instructions = pushInstruction(instructions, JNZBRACKET, o)
			}
		default:
			panic(fmt.Errorf("not added instruction - %s", i[0]))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("shouldn't see an error while loading bytecodes, %s", err))
	}

	return instructions
}

func pushInstruction(is []*Instruction, operator Mnemonic, operand interface{}) []*Instruction {
	i := &Instruction{Operator: operator, Operand: operand}
	is = append(is, i)
	return is
}

func (i *Instruction) String() string {
	switch i.Operator {
	case INC:
		return fmt.Sprintf("INC %d", i.Operand)
	case DEC:
		return fmt.Sprintf("DEC %d", i.Operand)
	case INCPTR:
		return fmt.Sprintf("INCPTR %d", i.Operand)
	case DECPTR:
		return fmt.Sprintf("DECPTR %d", i.Operand)
	case PUTCHAR:
		return "PUTCHAR"
	case GETCHAR:
		return "GETCHAR"
	case JZBRACKET:
		return fmt.Sprintf("JZBRACKET %d", i.Operand)
	case JNZBRACKET:
		return fmt.Sprintf("JNZBRACKET %d", i.Operand)
	default:
		panic(fmt.Errorf("not added instruction - %s", string(i.Operator)))
	}
}
