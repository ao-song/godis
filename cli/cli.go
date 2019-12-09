package main

import (
	"bufio"
	"fmt"
	"godis/lib"
	"os"
	"strings"
)

func main() {
	s := &kv.Store{KV: make(map[string]interface{})}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		runCmd(cmdStr, s)
	}
}

func runCmd(cmdStr string, store *kv.Store) {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	cmdStrList := strings.Fields(cmdStr)
	if len(cmdStrList) == 0 {
		return
	}
	switch cmdStrList[0] {
	case "get":
		if len(cmdStrList) != 2 {
			fmt.Print("Command incorrect!\n")
			return
		}
		v, found := store.Get(cmdStrList[1])
		if found {
			fmt.Print(v)
			fmt.Print("\n")
		} else {
			fmt.Print("Value not found!\n")
		}
	case "set":
		if len(cmdStrList) != 3 {
			fmt.Print("Command incorrect!\n")
			return
		}
		store.Set(cmdStrList[1], cmdStrList[2])
		fmt.Print("ok\n")
	case "del":
		if len(cmdStrList) != 2 {
			fmt.Print("Command incorrect!\n")
			return
		}
		store.Del(cmdStrList[1])
		fmt.Print("ok\n")
	case "exit":
		os.Exit(0)
	default:
		fmt.Print("Useage: set [key] [value] | get [key] | del [key]\n")
	}
}
