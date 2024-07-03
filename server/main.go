package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
	TYPE = "tcp"
)

type PairList struct {
	pairs map[uint32]int32
}

func main() {
	fmt.Println("Starting Prime Time Server...")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server listening on", HOST+":"+PORT)

	db := &PairList{
		pairs: make(map[uint32]int32),
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleRequest(conn, db)
	}
}

func handleRequest(conn net.Conn, db *PairList) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	buffer := make([]byte, 9)
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
		if n < 9 {
			fmt.Println("Less than 9 bytes in message")
			conn.Write([]byte("undefined"))
			continue
		}

		// Process the buffer to remove \n\r characters coming from telnet
		processedData := strings.ReplaceAll(string(buffer[:n]), "\r\n", "")
		fmt.Printf("Received: %s\n", processedData)

		if processedData[0] == 'I' {
			timestamp := binary.BigEndian.Uint32([]byte(processedData[1:5]))
			price := int32(binary.BigEndian.Uint32([]byte(processedData[5:9])))
			fmt.Printf("timestamp = %d, price = %d\n", timestamp, price)
			db.pairs[timestamp] = price
		} else if processedData[0] == 'Q' {
			minTime := binary.BigEndian.Uint32([]byte(processedData[1:5]))
			maxTime := binary.BigEndian.Uint32([]byte(processedData[5:9]))
			fmt.Printf("minTime = %d, maxTime = %d\n", minTime, maxTime)
			var mean int32 = 0
			var count int32 = 0
			for key, val := range db.pairs {
				if minTime <= key && key <= maxTime {
					mean += val
					count += 1
				}
			}
			if count > 0 {
				mean /= count
			}

			response := make([]byte, 4)
			binary.BigEndian.PutUint32(response, uint32(mean))
			conn.Write(response)
		} else {
			conn.Write([]byte("undefined"))
			continue
		}
	}
}
