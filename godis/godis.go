package main

import (
    "net"
    "log"
    "encoding/json"
    "godis/kv"
    "fmt"
    "io"
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

    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal("TCP server accept error:", err)
        }

        go handleConnection(conn, s)
    }
}

func handleConnection(conn net.Conn, s *kv.Store) {
    defer conn.Close()

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
        
        switch req.Action {
        case "set":
            setValue, ok := req.Value.(map[string]interface{})
            if !ok {
                log.Fatal("Set value not in correct format")
                break
            }
            for k, v := range setValue {
                s.Set(k, v)
            }
            encoder.Encode("ok")
        case "get":
            getKey, ok := req.Value.(string)
            if !ok {
                log.Fatal("Get key not in correct format")
                break
            }
            data, found := s.Get(getKey)
            if found {
                encoder.Encode(fmt.Sprintf(`{"value": %s}`, data))
            } else {
                encoder.Encode("Not found")
            }
        case "del":
            delKey, ok := req.Value.(string)
            if !ok {
                log.Fatal("Del key not in correct format")
                break
            }
            s.Del(delKey)
            encoder.Encode("ok")
        default:
            encoder.Encode("Incorrect command")
        }
    }    
}