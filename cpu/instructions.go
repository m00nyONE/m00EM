package cpu

const (
	NOP         = 0x0
	HLT         = 0xFF
	MOV_LIT_REG = 0x10
	MOV_REG_REG = 0x11
	MOV_REG_MEM = 0x12
	MOV_MEM_REG = 0x13
	ADD_REG_REG = 0x14
	JMP_EQ      = 0x20
	JMP_NOT_EQ  = 0x21
	JMP         = 0x2F
	PSH_LIT     = 0x30
	PSH_REG     = 0x31
	POP         = 0x3F
	CAL_LIT     = 0x40
	CAL_REG     = 0x41
	//JMP_GREATER_THEN = 0x99
	//JMP_LESS_THEN = 0x99
	//JMP_ZERO = 0x99
	//JMP_ONE = 0x99
	JMP_GREATER = 0x99
	JMP_LESS    = 0x99
	RET         = 0x4F
)
