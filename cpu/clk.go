package cpu

import (
	"time"
)

var ClkChan chan bool
var ClockSpeed time.Duration = 100 //in Hz

func Clk() {
	for {
		time.Sleep(time.Duration(int64(time.Second / ClockSpeed)))
		//fmt.Println("clk")
		ClkChan <- true
	}
}
