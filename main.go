package main

import (
	"log"

	s "github.com/Horryportier/AI_glove/server"
)

func main() {

        server := s.NewServer(":8090")

        go func() {
                for msg := range server.msgch {
                }
        }()
        log.Fatal(server.Start())
}
