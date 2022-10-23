package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"toks/serialDriver/serial"
)

func main() {
	portWrite := &serial.Port{Name: "/dev/ttys003", Baund: 9600}
	portRead := &serial.Port{Name: "/dev/ttys004", Baund: 9600}

	err := serial.InitPort(portRead)
	if err != nil {
		log.Fatal(err)
	}
	defer serial.Close(*portRead)

	err = serial.InitPort(portWrite)
	if err != nil {
		log.Fatal(err)
	}
	defer serial.Close(*portWrite)

	var data string
	fmt.Print("Input data: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	data = sc.Text()

	fmt.Println(serial.WritePackage(*portWrite, "", data))

	buf, err := serial.ReadPackage(*portRead)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read: ", buf)
}
