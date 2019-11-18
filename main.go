package main

import (
	"fmt"
	"github.com/m00nyONE/m00EM/cpu"
)

var Version string

//TODO: change everything to bytes instead of uint16

func init() {
	cpu.Version = Version
}

func main() {

	go cpu.Clk()

	for i := 0; <-cpu.ClkChan; i++ {
		cpu.Step()

		if cpu.DEBUG {
			fmt.Printf("Clock: %x\t", i)
			fmt.Printf("ACC: %d\n", cpu.GetValue("ACC"))
		}
	}
	defer close(cpu.ClkChan)
}
