package memory

var RAM []uint16

func Read(addr uint16) (data uint16) {
	return RAM[addr]
}
func Write(addr uint16, data uint16) {
	RAM[addr] = data
}

func Copy(src, drain uint16) {
	RAM[drain] = RAM[src]
}

func Move(src, drain uint16) {
	RAM[drain] = RAM[src]
	RAM[src] = 0x0
}

func Clear() {
	for i := 0; i < len(RAM); i++ {
		RAM[i] = 0x0
	}
}
