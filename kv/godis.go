package main

import (
    "net"
    "log"
    "encoding/json"
    "godis/kv"
)

type request struct {
    action string
    value  interface{}
}

type keyValue struct {
    key   string
    value string
}

func main() {
    s := &kv.Store{KV: make(map[string]interface{})}

    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal("TCP server listen error:", err)
    }

    defer ln.close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal("TCP server accept error:", err)
        }

        go handleConnection(conn, s)
    }
}

func handleConnection(conn net.Conn, s *kv.Store) {
    defer conn.close()

    decoder := json.NewDecoder(conn)

    for {
        var req request

        err := decoder.Decode(&req)
        if err != nil {
            log.Fatal("JSON decode error: ", err)

            if err == io.EOF {
                log.Print("Connection closed.")
                return
            }
        }
        
        switch req.action.(string) {
        case "set":           
            v := req.value

        case "get":
        case "del":
        default:
        }
    }    
}