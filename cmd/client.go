// socket-client project main.go
package cmd

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

type packet struct {
	id uint32
	t  time.Time
}

func Client(host string) error {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, host+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	c := make(chan packet, 3000)
	go readResponse(connection, c)
	token := make([]byte, 64*1024)
	rand.Read(token)
	log.Println(token)
	var i uint32 = 0

	for {
		binary.LittleEndian.PutUint32(token, i)
		///send some data
		_, err = connection.Write(token)
		log.Println("Send " + fmt.Sprint(i))
		c <- packet{i, time.Now()}
		if err != nil {
			log.Println("Error reading:", err.Error())
		}
		i += 1
	}
}

func readResponse(connection net.Conn, c chan packet) {
	buffer := make([]byte, 4)
	for {
		// mLen, err := io.ReadAtLeast(connection, buffer, len(buffer))
		p := <-c
		_, err := io.ReadAtLeast(connection, buffer, 4)
		if err != nil {
			log.Println("Error reading:", err.Error())
		}
		id := binary.LittleEndian.Uint32(buffer[:4])

		if id != p.id {
			log.Println("Error id: ", id, " ", p.id)
		}

		elp := time.Now().Sub(p.t).Milliseconds()

		log.Println("Received: ", fmt.Sprint(binary.LittleEndian.Uint32(buffer[:4])), " ", elp)

		if elp > 50 {
			log.Println("WARN Received: ", fmt.Sprint(binary.LittleEndian.Uint32(buffer[:4])), " ", elp)
		}

		// log.Println("Received: ", fmt.Sprint(binary.LittleEndian.Uint32(buffer[:4])), " ", mLen)
	}
}
