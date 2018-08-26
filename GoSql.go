package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	Entries []DbEntry
}

type DbEntry struct {
	Key   string
	Value string
}

var (
	db Database
)

func main() {

	db = getDatabase()

	//Open Port
	port := 3333
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

	//Connection error handing
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}

	//Begin Listening
	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go processConnection(conn)
	}

}

func processConnection(conn net.Conn) {

	defer saveDatabase()
	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		switch err {
		case io.EOF:
			break
		case nil:
			log.Println("Receive:", data)
			if isTransportOver(data) {
				parsedCommand := strings.Fields(data)
				command := parsedCommand[0]
				switch command {
				case "CREATE":
					key := parsedCommand[1]
					value := parsedCommand[2]
					handleCreate(key, value)
					w.Write([]byte("CREATE SUCCESSFUL"))
					break
				case "GET":
					key := parsedCommand[1]
					result := handleGet(key)
					w.Write([]byte(result))
					break
				case "UPDATE":
					key := parsedCommand[1]
					value := parsedCommand[2]
					handleUpdate(key, value)
					w.Write([]byte("UPDATE SUCCESSFUL"))
					break
				case "DELETE":
					key := parsedCommand[1]
					handleDelete(key)
					w.Write([]byte("DELETE SUCCESSFUL"))
					break
				case "SAVE":
					saveDatabase()
					w.Write([]byte("SAVE SUCCESSFUL"))
				default:
					w.Write([]byte("Error: Command Not Found"))
				}
				break
			}

		default:
			log.Fatalf("Receive data failed:%s", err)
			w.Write([]byte("failure"))
			return
		}
		break
	}
	w.Flush()
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

//This function pulls all the data out of the JSON file. It will be inserted into the DB object and manipulated
func getDatabase() (db Database) {
	raw, err := ioutil.ReadFile("database.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	json.Unmarshal(raw, &db)
	return db
}

//This function will be called at the end of the program to update our JSON database file
func saveDatabase() {
	//Create output JSON
	bytes, err := json.Marshal(db)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//Open Output file
	fo, err := os.Create("database.json")
	if err != nil {
		panic(err)
	}

	//Write to File
	if _, err := fo.Write(bytes); err != nil {
		panic(err)
	}
}

//Command Functions
func handleCreate(key string, value string) {
	newEntry := DbEntry{Key: key, Value: value}
	db.Entries = append(db.Entries, newEntry)
}

func handleGet(key string) (result string) {
	for _, v := range db.Entries {
		if v.Key == key {
			result = v.Value
		}
	}
	return result
}

func handleUpdate(key string, value string) {
	updatedDB := new(Database)
	for _, v := range db.Entries {
		if v.Key != key {
			newEntry := DbEntry{Key: v.Key, Value: v.Value}
			updatedDB.Entries = append(updatedDB.Entries, newEntry)
		} else {
			newEntry := DbEntry{Key: key, Value: value}
			updatedDB.Entries = append(updatedDB.Entries, newEntry)
		}
	}

	db = *updatedDB
}

func handleDelete(key string) {
	updatedDB := new(Database)
	for _, v := range db.Entries {
		if v.Key != key {
			newEntry := DbEntry{Key: v.Key, Value: v.Value}
			updatedDB.Entries = append(updatedDB.Entries, newEntry)
		}
	}

	db = *updatedDB
}
