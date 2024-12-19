package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/digitalcrab/adventofcode/utils"
)

type Machine struct {
	A, B, C int
}

func (m Machine) String() string {
	return fmt.Sprintf("A: %d, B: %d, C: %d", m.A, m.B, m.C)
}

func (m Machine) Run(programm []int) (out []int) {
	var pointer int

	for pointer < len(programm)-1 { // if there are at least 2 elements
		// first and second are opcode and operand
		opcode := programm[pointer]
		operand := programm[pointer+1]
		// how much increment the pointer
		// we read 2 instructions (there is an exception later)
		inc := 2

		switch opcode {
		case 0: // adv
			// The numerator is the value in the A register.
			// The denominator is found by raising 2 to the power of the instruction's combo operand.
			//
			// note for myself:
			// left-shifting an integer “a” with an integer “b” denoted as ‘(a<<b)’
			// is equivalent to multiplying a with 2^b (2 raised to power b)
			// so
			// denominator := 1 << m.combo(operand)
			//
			// right-shifting an integer “a” with an integer “b” denoted as ‘(a>>b)‘
			// is equivalent to dividing a with 2^b.
			// that is exactly what we need here:
			// register A dividing with 2 to the power of combo operand

			// The result of the division operation is truncated to an integer and then written to the A register.
			m.A = m.A >> m.combo(operand)

		case 1: // bxl
			// calculates the bitwise XOR of register B and the instruction's literal operand,
			// then stores the result in register B.
			m.B = m.B ^ operand

		case 2: // bst
			// calculates the value of its combo operand modulo 8, then writes that value to the B register.
			m.B = m.combo(operand) % 8

		case 3: // jnz
			// does nothing if the A register is 0.
			// if the A register is not zero, it jumps by setting the instruction pointer to the value
			// of its literal operand. if this instruction jumps, the instruction pointer is not increased
			// by 2 after this instruction.
			if m.A > 0 {
				pointer = operand
				inc = 0
			}

		case 4: // bxc
			// calculates the bitwise XOR of register B and register C,
			// then stores the result in register B.
			// this instruction reads an operand but ignores it.
			m.B = m.B ^ m.C

		case 5: // out
			// calculates the value of its combo operand modulo 8, then outputs that value
			out = append(out, m.combo(operand)%8)

		case 6: // bdv
			// works exactly like the adv instruction except that the result is stored in the B register.
			// (The numerator is still read from the A register.)
			m.B = m.A >> m.combo(operand)

		case 7: // cdv
			// works exactly like the adv instruction except that the result is stored in the C register.
			// (The numerator is still read from the A register.)
			m.C = m.A >> m.combo(operand)

		default:
			panic(fmt.Sprintf("unexprect opcode: %d", opcode))
		}

		pointer += inc
	}

	return
}

func (m Machine) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return m.A
	case 5:
		return m.B
	case 6:
		return m.C
	default:
		panic(fmt.Sprintf("unexprect operand: %d", operand))
	}
}

// 2,4,1,1,7,5,1,5,0,3,4,4,5,5,3,0
//
// 0)
//    A = 23999685
// 1) ==== 2,4
//    B <- A % 8       (0b001 011 011 100 011 010 011 000 101 & 0b111 = 0b101) // looks like the intention is to take last 3 bits
//    B = 5
//    as it takes last 3 bits, we get always from 0 to 7
// 2) ==== 1,1
//    B <- B ^ 1       (0b101 ^ 0b001 = 0b100) // basically flips the last bit
//    B = 4
//    as we flip last bit from 0-7, we basically get back same 0-7 just in a different order [1,0,3,2,5,4,7,6]
// 3) ==== 7,5
//    C <- A >> B      (0b001 011 011 100 011 010 011 000 101 >> 4 = 0b101 101 110 001 101 001 100)
//    C = 1499980
//    as we move to the right, we chop up to 7 least significant bits from A
// 4) ==== 1,5
//    B <- B ^ 5       (0b100 ^ 0b101 = 0b001) // flips first and last bit
//    B = 1
// 5) ==== 0,3
//    A <- A >> 3      (0b001 011 011 100 011 010 011 000 101 >> 3 = 0b001 011 011 100 011 010 011 000)
//    A = 2999960
// 6) ==== 4,4
//    B <- B ^ C       (0b001 ^ 0b101 101 110 001 101 001 100 = 0b101 101 110 001 101 001 101) // flips the last bit
//    B = 1499981
// 7) ==== 5,5
//    out <- B % 8     (0b101 101 110 001 101 001 101 & 0b111 = 0b101)
//    we output only last 3 bits, so basically 0-7
// 8) ==== 3,0
//    if A > 0 jump to the beginning
//
// HINTS:
// & - 1 only when both are 1
// ^ - 1 only if bits are different
//
// Modulo:
// When the divisor is a power of 2, the modulo operation can be directly tied to
// the lower bits of the number. This makes it very efficient at the bit level.
// A % 8 = A & (8 - 1)
// 8 - 1 = 7 (0b111)
// 23999685 = 0b001 011 011 100 011 010 011 000 101
//
// 0b001 011 011 100 011 010 011 000 101 - 23999685
// 0b000 000 000 000 000 000 000 000 111 - 7
// 0b000 000 000 000 000 000 000 000 101 - 5

// FindBackwards finds correct register A
//
//	function is hardcoded by my input and does not work for every given input
func FindBackwards(programm []int, answer int) int {
	// exit
	if len(programm) == 0 {
		return answer
	}

	lastIdx := len(programm) - 1

	// as we output only last 3 bits of the number, we should also start from them
	for n := range slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7}) {
		a := (answer << 3) + n // backward of step 5
		b := a % 8             // step 1
		b = b ^ 1              // step 2
		c := a >> b            // step 3
		b = b ^ 5              // step 4
		b = b ^ c              // step 6
		out := b % 8           // step 7
		if out == programm[lastIdx] {
			next := FindBackwards(programm[:lastIdx], a)
			if next == -1 { // for that digit (0-7) not found lets try another one
				continue
			}
			return next
		}
	}

	// no solution found
	return -1
}

//go:embed "example.txt"
var exampleInput string

//go:embed "input.txt"
var DayInput string

func main() {
	machine, programm := Parse(exampleInput)
	fmt.Printf("Machine:\n%s\nProgramm:\n%v\n", machine, programm)
	out := machine.Run(programm)
	fmt.Printf("Output:\n%s\n", strings.Join(utils.IntsToStrings(out), ","))

	_, dayProgramm := Parse(DayInput)
	correctA := FindBackwards(dayProgramm, 0) // 0 is the last digit from the output
	fmt.Printf("Correct A: %d\n", correctA)
}
