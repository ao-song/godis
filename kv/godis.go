package main

import (
	"net"
	"log"
	"encoding/json"
	"godis/kv"
)

type request struct {
	action string
	value  string
}

func main() {
	s := &kv.Store{KV: make(map[string]interface{})}

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
        log.Fatal("TCP server listen error:", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("TCP server accept error:", err)
		}

		go handleConnection(conn, s)
	}
}

func handleConnection(conn net.Conn, s *kv.Store) {
	decoder := json.NewDecoder(conn)

	for {
		var req request

		err := decoder.Decode(&req)
		if err != nil {
			log.Fatal("JSON decode error:", err)
		}

		
	}
	
	
}

