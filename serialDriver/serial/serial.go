package serial

import (
	"time"
	"toks/serialDriver/netPackage"

	"github.com/tarm/serial"
)

type Port struct {
	Name       string
	Baund      int
	SerialPort *serial.Port
}

func ReadByte(p Port) (int, byte, error) {
	buf := make([]byte, 1)
	var n int
	var err error
	go func() { n, err = p.SerialPort.Read(buf) }()
	time.Sleep(100 * time.Microsecond)
	if err != nil {
		return 0, 0, err
	}
	return n, buf[0], nil

}

func ReadPackage(p Port) (string, error) {
	var newPackFlag, escFlag bool
	var data []byte

	for {
		n, b, err := ReadByte(p)
		if err != nil {
			return "", err
		}
		if n == 0 {
			break
		}
		// fmt.Println(string(b), newPackFlag, escFlag)
		if b == byte(netPackage.Flag) && !escFlag && !newPackFlag {
			newPackFlag = true
			continue
		}
		if b == byte(netPackage.Esc) && !escFlag {
			escFlag = true
			continue
		}
		if b == byte(netPackage.Flag) && escFlag {
			escFlag = false
		}

		data = append(data, b)

	}

	return string(data), nil
}

func WriteByte(p Port, text []byte) (int, error) {
	n, err := p.SerialPort.Write(text)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func WritePackage(p Port, address, text string) (int, error) {
	n, err := p.SerialPort.Write(netPackage.CreatePackage(address, text))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func InitPort(p *Port) error {
	c := &serial.Config{Name: p.Name, Baud: p.Baund}
	temp, err := serial.OpenPort(c)
	p.SerialPort = temp
	return err
}

func Close(p Port) error {
	if err := p.SerialPort.Close(); err != nil {
		return err
	}
	return nil
}

func ChangeSpeed(p *Port) error {

	err := InitPort(p)
	if err != nil {
		return err
	}
	return nil
}
