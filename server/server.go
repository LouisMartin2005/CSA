package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
	println("ERRoR you gooner")
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.

	//ln, _ = net.Listen("tcp", ":8030")
	for {
		conn, _ := ln.Accept()
		conns <- conn
	}

}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.

	reader := bufio.NewReader(client) // to declare a new reader for that connection
	for {
		msg, _ := reader.ReadString('\n') // reads input from a buffered reader until it hits a newline character
		msgs <- Message{clientid, msg}    // need to receive a channel in the type of Message which is (int,string)
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse() //read the command line argument and set the flag values to change the values in above if not default

	//TODO Create a Listener for TCP connections on the port given above.

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	ln, _ := net.Listen("tcp", *portPtr)
	ids := 0

	go acceptConns(ln, conns)

	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			// - start to asynchronously handle messages from this client
			ids++
			id := ids
			clients[id] = conn
			go handleClient(conn, id, msgs)

		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for id, clientConn := range clients {
				if id != msg.sender {
					fmt.Fprintf(clientConn, msg.message)
				}
			}
		}
	}
}
