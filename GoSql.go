package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	//Import the Database from the json file
	//getDatabase()

	//Run the service/listener
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func getDatabase(message string) {
	//This will return the json file that is the database. The queries will happen against the object that is returned
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime)) // don't care about return value
	// we're finished with this client
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
		//Need a function here to write the database back to the file
	}
}
