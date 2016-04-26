package main

import (
	"fmt"
	"log"

	"github.com/david-fry/go-3dprinter"
)

func main() {
	fmt.Println("Connecting to Printer...")
	p := printer.Connect("COM3", 115200)

	fmt.Println("Moving Extruder...")
	err := p.SendCommand("G28")
	err = p.SendCommand("G1 Z20")
	err = p.SendCommand("G1 Y100")
	err = p.SendCommand("G1 X100")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Heating Extruder...")
	err = p.SendCommand("M109 S100")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cooling Extruder...")
	err = p.SendCommand("M109 S100")

	if err != nil {
		log.Fatal(err)
	}

}
