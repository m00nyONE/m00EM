package cpu

import (
	"flag"
	"fmt"
	"github.com/m00nyONE/m00EM/examples"
	"github.com/m00nyONE/m00EM/memory"
	"log"
	"os"
	"time"
)

var versionP = flag.Bool("version", false, "display current version")
var debugP = flag.Bool("debug", false, "enables debugmode")
var clockSpeedP = flag.Int64("clockspeed", 1, "clockspeed in Hz")
var ramSizeP = flag.Int64("ramsize", 0xFFFF, "Size of RAM (min 0xFF) - (max 0xFFFF)")
var loadP = flag.String("load", "", "path of the program to load") // will later be replaced by a path to a .bin file

func init() {

	/*
		Flags
	*/
	flag.Parse()
	if *versionP {
		fmt.Println("Version: " + Version)
		os.Exit(0)
	}

	if *debugP {
		DEBUG = true
	}
	if *clockSpeedP <= 0 {
		log.Fatal("Clockspeed must be at least 1 Hz")
	} else {
		ClockSpeed = time.Duration(*clockSpeedP)
	}
	if *ramSizeP >= 0xFFF && *ramSizeP <= 0xFFFF {
		memory.RAM = make([]uint16, *ramSizeP)
	} else {
		log.Fatal("Ram Size is not within bounds - min 0xFFF -- max 0xFFFF")
	}
	if len(*loadP) == 0 {
		examples.Fibonacci()
	} else {
		//TODO: Load file and read it
	}

	/*
		Create registers & clear them
	*/
	register = make(map[string]uint16)
	for i := 0; i < len(registerNames); i++ {
		register[registerNames[i]] = 0x0
	}

	/*
		Setup VRAM
		TODO: i did't decide how this will work for now.
		There are two options:
			1. extra GPU with extra instructions and an instruction for the CPU to send it to the GPU
			2. 16x64 ascii console ( last 1024 addresses on RAM ). Something like in the Intel 6068

			the first one would be more fun i think
	*/

	/*
		set StackPointer to last RAM adress
	*/
	register["SP"] = uint16(len(memory.RAM) - 1)

	/*
		Set up Clk Channel
	*/
	ClkChan = make(chan bool)
}
