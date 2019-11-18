package cpu

import "log"

var stackSize uint16 = 0
var register map[string]uint16
var registerNames = []string{
	"IP",  //Instruction Pointer
	"FP",  //Frame Pointer
	"ACC", //Accumulator
	"AX",  //Address Register
	"CX",  //Counter Register
	"IR",  //Instruction Register
	"ICR", //Interrupt Register
	"SP",  //Stack Pointer
	"R1",  //Register1
	"R2",  //Register2
	"R3",  //Register3
	"R4",  //Register4
	"SR",  //Status Register
}

/*
	stackPointer functions
*/
func Increment(reg string) {
	register[reg] += 1
}
func Decrement(reg string) {
	register[reg] -= 1
}
func SetValue(reg string, value uint16) {
	register[reg] = value
}
func GetValue(reg string) (value uint16) {
	if val, ok := register[reg]; ok {
		return val
	} else {
		log.Fatal("Register not Found! Check Code...")
		return
	}
}
func ShiftRight(reg string, n int) {
	for i := 0; i < n; i++ {
		register[reg] = register[reg] >> 1
	}
}
func ShiftLeft(reg string, n int) {
	for i := 0; i < n; i++ {
		register[reg] = register[reg] << 1
	}
}
func RotateRight(reg string, n int) {
	register[reg] = (register[reg] >> n) | (register[reg] << (16 - n))
}
func RotateLeft(reg string, n int) {
	s := uint(n) & (16 - 1)
	register[reg] = register[reg]<<s | register[reg]>>(16-s)
}
func Negate(reg string) {
	register[reg] = ^register[reg]
}
func TwosComplement(reg string) {
	register[reg] = ^register[reg] + 1
}
