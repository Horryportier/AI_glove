package main

import (
	"fmt"
	"log"

	server "github.com/Horryportier/AI_glove/server"
)

func main() {

	s := server.NewServer(":8090")

	go func() {
			for msg := range s.Msgch {
				fmt.Printf("msg from conn (%v): %s\n",
					server.GoodStyle.Render(msg.From), string(msg.Payload))

			}
	}()

	log.Fatal(s.Start())
}
