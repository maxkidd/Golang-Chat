package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)
var conns []net.Conn

type Message struct{
	sender string
	message string
}

var messageMutex sync.Mutex
var messages []Message
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	go writeConnections()
	for
	{
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conns = append(conns, c)
		go handleConnection(c)
	}
}

func writeConnections() {
	for{
		if len(messages) > 0{
			messageMutex.Lock()
			for _, c := range conns{
				for _, m := range messages{
					if c.RemoteAddr().String() != m.sender{
						c.Write([]byte(m.message))
					}
				}
			}
			messages = nil
			messageMutex.Unlock()
		}
		time.Sleep(1* time.Second)
	}
}



func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Print("Received message: ", string(netData))
		messageMutex.Lock()
		messages = append(messages,Message{message: netData, sender: c.RemoteAddr().String()})
		messageMutex.Unlock()
	}
	c.Close()
}