package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {

	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	textData := []byte("I")
	intData := make([]byte, 4)
	int2Data := make([]byte, 4)
	binary.BigEndian.PutUint32(intData, uint32(12345))
	binary.BigEndian.PutUint32(int2Data, uint32(101))

	// Send the data
	_, err = conn.Write(textData)
	if err != nil {
		fmt.Println("Error sending text data:", err)
		return
	}

	_, err = conn.Write(intData)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	_, err = conn.Write(int2Data)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	binary.BigEndian.PutUint32(intData, uint32(12346))
	binary.BigEndian.PutUint32(int2Data, uint32(102))

	// Send the data
	_, err = conn.Write(textData)
	if err != nil {
		fmt.Println("Error sending text data:", err)
		return
	}

	_, err = conn.Write(intData)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	_, err = conn.Write(int2Data)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	binary.BigEndian.PutUint32(intData, uint32(12347))
	// var negative int32 = -100
	binary.BigEndian.PutUint32(int2Data, uint32(100))

	// Send the data
	_, err = conn.Write(textData)
	if err != nil {
		fmt.Println("Error sending text data:", err)
		return
	}

	_, err = conn.Write(intData)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	_, err = conn.Write(int2Data)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	binary.BigEndian.PutUint32(intData, uint32(40960))
	binary.BigEndian.PutUint32(int2Data, uint32(5))

	// Send the data
	_, err = conn.Write(textData)
	if err != nil {
		fmt.Println("Error sending text data:", err)
		return
	}

	_, err = conn.Write(intData)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	_, err = conn.Write(int2Data)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	textData = []byte("Q")
	binary.BigEndian.PutUint32(intData, uint32(12288))
	binary.BigEndian.PutUint32(int2Data, uint32(16384))

	// Send the data
	_, err = conn.Write(textData)
	if err != nil {
		fmt.Println("Error sending text data:", err)
		return
	}

	_, err = conn.Write(intData)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	_, err = conn.Write(int2Data)
	if err != nil {
		fmt.Println("Error sending int32 data:", err)
		return
	}

	result := make([]byte, 4)
	conn.Read(result)
	value := binary.BigEndian.Uint32(result)
	fmt.Println("Value:", value)
}
