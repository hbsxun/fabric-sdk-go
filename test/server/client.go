package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Model struct {
	Owner  string
	Name   string
	Source string
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Dial error %v\n", err)
		os.Exit(-1)
	}
	model := &Model{"Alice", "M1", "Software..."}
	modelByte, err := json.Marshal(model)
	if err != nil {
		fmt.Printf("model Marshal return err [%s]\n", err.Error())
		return
	}
	if _, err = conn.Write(modelByte); err != nil {
		fmt.Printf("conn Write return err [%s]\n", err.Error())
		return
	}
	buf := make([]byte, 512)
	size, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn Read return err [%s]\n", err.Error())
		return
	}
	fmt.Printf("Client Receive: [%s]\n", string(buf[:size]))
}
