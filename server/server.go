package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Payload struct {
	Hall int
	Mpu  Mpu
}

type Mpu struct {
	Status       string
	Acceleration Acceleration
	Rotation     Rotation
	Temp         float64
}

type Acceleration struct {
	X float64
	Y float64
	Z float64
}
type Rotation struct {
	X float64
	Y float64
	Z float64
}

func Connect(SERVER_HOST, SERVER_PORT, SERVER_TYPE string) error {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		return err
	}

	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			print("help!!")
			return err
		}
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	var p Payload
	payload := string(buffer[:mLen])
	p, err = parseJson(payload, p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Mpu\nAccereration: (%v,%v,%v)\nRotation: (%v,%v,%v)\nTemp: (%v)\nHall: %v\n",
		p.Mpu.Acceleration.X, p.Mpu.Acceleration.Y, p.Mpu.Acceleration.Z,
		p.Mpu.Rotation.X, p.Mpu.Rotation.Y, p.Mpu.Rotation.Z, p.Mpu.Temp, p.Hall)

	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}

func parseJson(payload string, p Payload) (Payload, error) {

	err := json.Unmarshal([]byte(payload), &p)
	if err != nil {
		return p, err
	}

	return p, nil
}
