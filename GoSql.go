package main

import (
	"fmt"
	"net"
	"os"
)

var commandHandlers = map[string]interface{}{
	"PUT":       handlePut,
	"GET":       handleGet,
	"GETLIST":   handleGetlist,
	"PUTLIST":   handlePutlist,
	"INCREMENT": handleIncrement,
	"APPEND":    handleAppend,
	"DELETE":    handleDelete,
	"STATS":     handleStats,
}

type message struct {
	command   string
	key       string
	value     string
	valueType string
}

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
		go parseMessage(conn)
	}
}

//Database Helper Functions
func getDatabase() {
	//This will return the json file that is the database. The queries will happen against the object that is returned
}

func parseMessage(conn net.Conn) {
	defer conn.Close()
	//var mess message
	//mess = conn.Read
	//fmt.Println(mess)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
		//Need a function here to write the database back to the file
	}
}

//Command Functions
func handlePut() {

}

func handleGet() {

}

func handleGetlist() {

}

func handlePutlist() {

}

func handleIncrement() {

}

func handleAppend() {

}

func handleDelete() {

}

func handleStats() {

}
