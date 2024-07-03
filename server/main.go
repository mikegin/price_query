package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	HOST         = "0.0.0.0"
	PORT         = "8080"
	TYPE         = "tcp"
	MESSAGE_SIZE = 9
)

func main() {
	fmt.Println("Starting Price Query Server...")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server listening on", HOST+":"+PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	buffer := make([]byte, MESSAGE_SIZE)

	pairs := map[int32]int32{}

	for {
		n, err := io.ReadFull(reader, buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from connection:", err)
			}
			conn.Write([]byte("undefined"))
			conn.Close()
			return
		}
		if n != MESSAGE_SIZE {
			fmt.Println("Message isn't 9 bytes")
			conn.Write([]byte("undefined"))
			continue
		}

		if buffer[0] == 'I' {
			var timestamp int32
			var price int32
			binary.Read(bytes.NewBuffer(buffer[1:5]), binary.BigEndian, &timestamp)
			binary.Read(bytes.NewBuffer(buffer[5:9]), binary.BigEndian, &price)
			fmt.Printf("timestamp = %d, price = %d\n", timestamp, price)
			pairs[timestamp] = price
		} else if buffer[0] == 'Q' {
			var minTime int32
			var maxTime int32
			binary.Read(bytes.NewBuffer(buffer[1:5]), binary.BigEndian, &minTime)
			binary.Read(bytes.NewBuffer(buffer[5:9]), binary.BigEndian, &maxTime)
			fmt.Printf("minTime = %d, maxTime = %d\n", minTime, maxTime)
			// use int64 since working with large values like 1178774581940 -- FAIL:Q 285864834 377826687: expected 49374825 (1178774581940/23874), got 81827
			var mean int64 = 0
			var count int64 = 0
			for key, val := range pairs {
				if minTime <= key && key <= maxTime {
					mean += int64(val)
					count += 1
				}
			}
			if count > 0 {
				mean /= count
			}

			response := &bytes.Buffer{}
			binary.Write(response, binary.BigEndian, int32(mean))
			io.Copy(conn, response)
		} else {
			fmt.Println("Illegal message type received: ", buffer)
			conn.Write([]byte("undefined"))
			continue
		}
	}
}
