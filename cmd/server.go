// socket-server project main.go
package cmd

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "20000"
	SERVER_TYPE = "tcp"
)

func Server() error {
	log.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	log.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	log.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		log.Println("client connected")
		go processClient(connection)
	}
}
func processClient(connection net.Conn) {
	buffer := make([]byte, 64*1024)
	defer connection.Close()
	for {
		// mLen, err := connection.Read(buffer)
		start := time.Now()
		mLen, err := io.ReadAtLeast(connection, buffer, len(buffer))
		if err != nil {
			log.Println("Error reading:", err.Error())
			return
		}
		end := time.Now()

		elp := end.Sub(start)

		id := binary.LittleEndian.Uint32(buffer)
		log.Println("Received: ", id, " ", mLen)

		if elp > 50 {
			log.Println("WARN: Received: ", id, " ", elp)
		}
		// log.Println("Received: ", string(buffer[:mLen]))
		token := make([]byte, 4)
		binary.LittleEndian.PutUint32(token, id)
		_, err = connection.Write(token)
	}
}
