package main

import (
	"log"
	"github.com/Horryportier/AI_glove/server"
)

var connections int = 1

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "8090"
	SERVER_TYPE = "tcp"
)


func main() {
        err := server.Connect(SERVER_HOST, SERVER_PORT, SERVER_TYPE)
        if err != nil {
                log.Fatal(err)
        }
}


