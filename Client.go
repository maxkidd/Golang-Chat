package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Require host:port argument")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	go handleMessages(c)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")


	}
	fmt.Println("hello world")
}

func handleMessages(c net.Conn) {
	for {
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("R ->: " + message)
	}
}
