package main

import (
	"bufio"
	//"flag"
	"fmt"
	"net"
	"os" // a better way than flag
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	//msgP := flag.String("msg", "Default message", "The message you wanted to send") // a pointer to a string
	////msg := "hello from client"
	//flag.Parse() // to be able to type stuff from the terminal (go run client.go -msg "new message from louis")
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

	for {
		fmt.Printf("Enter Text ->")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		read(conn)
	}
}
