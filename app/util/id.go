package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	uuidString := id.String()
	fmt.Println(uuidString)

	//uuid.NewUUID()
}
