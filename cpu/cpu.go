package cpu

import (
	"github.com/m00nyONE/m00EM/memory"
	"log"
)

var DEBUG = false
var Version string

func Reset() {
	memory.Clear()
}

func fetch() (instruction uint16) {
	instruction = memory.Read(GetValue("IP"))
	Increment("IP")
	return instruction

}

//maybe used in the future to only fetch single bytes and not always uint16
func fetch16() {

}

func Push(value uint16) {
	memory.Write(GetValue("SP"), value)
	SetValue("SP", GetValue("SP")-1)
	stackSize += 1
}
func Pop() (value uint16) {
	SetValue("SP", GetValue("SP")+1)
	stackSize -= 1
	return memory.Read(GetValue("SP"))
}

func PushState() {
	Push(GetValue("R1"))
	Push(GetValue("R2"))
	Push(GetValue("R3"))
	Push(GetValue("R4"))
	Push(GetValue("IP"))
	Push(stackSize + 1)
	SetValue("FP", GetValue("SP"))
	stackSize = 0
}

func PopState() {
	SetValue("SP", GetValue("FP"))
	stackSize = Pop()
	SetValue("IP", Pop())
	SetValue("R4", Pop())
	SetValue("R3", Pop())
	SetValue("R2", Pop())
	SetValue("R1", Pop())
	nArgs := Pop()
	for i := 1; i <= int(nArgs); i++ {
		Pop()
	}
	SetValue("FP", GetValue("FP")+stackSize)
}

func execute(instruction uint16) {
	switch instruction {

	case NOP:
		return
	case HLT:
		log.Printf("HLT (%x) called -- Exit", instruction)
		return
	case MOV_LIT_REG:
		var literal = fetch()
		var register = fetch()
		SetValue(registerNames[register], literal)
		return
	case MOV_REG_REG:
		var source = fetch()
		var drain = fetch()
		SetValue(registerNames[drain], GetValue(registerNames[source]))
		return
	case MOV_REG_MEM:
		var register = fetch()
		var address = fetch()
		memory.Write(address, GetValue(registerNames[register]))
		return
	case MOV_MEM_REG:
		var address = fetch()
		var register = fetch()
		SetValue(registerNames[register], memory.Read(address))
		return
	case ADD_REG_REG:
		reg1 := fetch()
		reg2 := fetch()
		SetValue("ACC", GetValue(registerNames[reg1])+GetValue(registerNames[reg2]))
		return
	case JMP_EQ:
		value := fetch()
		address := fetch()
		if value == GetValue("ACC") {
			SetValue("IP", address)
		}
		return
	case JMP_NOT_EQ:
		value := fetch()
		address := fetch()
		if value != GetValue("ACC") {
			SetValue("IP", address)
		}
		return
	case JMP:
		address := fetch()
		SetValue("IP", address)
		return
	case PSH_LIT:
		Push(fetch())
		return
	case PSH_REG:
		Push(GetValue(registerNames[fetch()]))
		return
	case POP:
		register := fetch()
		SetValue(registerNames[register], Pop())
		return
	case CAL_LIT:
		address := fetch()
		PushState()
		SetValue("IP", address)
		return
	case CAL_REG:
		register := fetch()
		value := GetValue(registerNames[register])
		PushState()
		SetValue("IP", value)
		return
	case RET:
		PopState()
		return
	default:
		log.Fatalf("Can not execute instruction: %x --- EXIT", instruction)
	}
}

func Step() {
	execute(fetch())
}
