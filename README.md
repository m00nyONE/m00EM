# m00EM 
## an emulated CPU in Go

This in a very early state - Changes in the opcode are most likely to happen.

### Instructions:

- i know that source & drain are currently not in the right order. this will be fixed later

| Operation | Instruction | arg1 | arg2 | description|
|---|---|---|---|---|
| NOP | 0x00 | - | - | No operation|
| HLT | 0xff | - | - | Stops the CPU|
| MOV_LIT_REG | 0x10 | Literal | Register | moves a literal into a register|
| MOV_REG_REG | 0x11 | From Register | To Register | moves a register into a register|
| MOV_REG_MEM | 0x12 | From Register | To Address | moves a register into Memory|
| MOV_MEM_REG | 0x13 | From Memory | To Register | moves Memory into a register|
| ADD_REG_REG | 0x14 | Register | Register | adds registers & saves result to ACC|
| JMP_EQ | 0x20 | Literal | Address | compares ACC & Literal - jump to address if equal|
| JMP_NOT_EQ | 0x21 | Literal | Address | compares ACC & Literal - jump to address if not equal|
| JMP | 0x2F | Address | - | compares ACC & Literal - jump if not equal|
| PSH_LIT | 0x30 | Literal | - | push Literal onto stack|
| PSH_REG | 0x31 | Register | - | push Register onto stack|
| POP | 0x3F | Register | - | pop value from stack into Register|
| CAL_LIT | 0x40 | Address | - | jumps to subroutine at Address|
| CAL_REG | 0x41 | Address | - | jumps to subroutine at Address given by the value of the Register|
| RET | 0x4F | - | - | returns from a subroutine and restores the previous state|

### Parameters

| flag | value | default | description |
|---|---|---|---|
| -debug| bool | false | enables debugmode |
| -clockspeed| Int | 1 | clockspeed in Hz |
| -ramsize| Int | 0xFFFF | Size of RAM ( 0xFFF - 0xFFFF ) |
| -load| string | "" | path of the program to load !!! currently disabled  |
