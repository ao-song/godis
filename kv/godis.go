package main

import (
    "net"
    "log"
    "encoding/json"
    "godis/kv"
)

type request struct {
    Action string
    Value  interface{}
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
    encoder := json.NewEncoder(conn)

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
        
        switch req.Action.(string) {
        case "set":
            for k, v := range req.Value {
                s.Set(k, v)
            }
            encoder.Encode("ok")
        case "get":
            data, found := store.Get(req.Value)
            if found {
                encoder.Encode(fmt.Sprintf(`{"value": %s}`, data))
            } else {
                encoder.Encode("Not found")
            }
        case "del":
            s.Del(req.Value)
            encoder.Encode("ok")
        default:
            encoder.Encode("Incorrect command")
        }
    }    
}