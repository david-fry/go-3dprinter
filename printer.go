package printer

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

//Printer is a connection to a reprap or similar printer
type Printer struct {
	s *serial.Port
}

//Connect creates the printer struct and returns it after initing it
func Connect(port string, speed int64) Printer {
	c := &serial.Config{Name: "COM3", Baud: 115200, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	p := Printer{}

	p.s = s

	p.readPump()
	p.readPump()

	return p
}

func (p *Printer) readPump() string {
	output := ""
	oldLength := -1
	for {
		buf := make([]byte, 128)
		n, err := p.s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		output += fmt.Sprintf("%s", buf[:n])
		if len(output) == oldLength {
			return output
		}

		oldLength = len(output)

	}
}

//SendCommand sends a single GCODE command to the printer
func (p *Printer) SendCommand(g string) error {
	g = g + "\n"
	_, err := p.s.Write([]byte(g))
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasSuffix(p.readPump(), "ok\n") {
		return errors.New("command did not complete successfully")
	}

	return nil
}
