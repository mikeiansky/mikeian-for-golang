package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	uuidString := uuid.String()
	fmt.Println(uuidString)
}
