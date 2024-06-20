package main

import (
	"log"

	"github.com/pedrobarbosak/go-errors"
	"github.com/pedrobarbosak/go-errors/example/service"
)

func main() {
	err := errors.New("this is a example error with some variables:", true, 1)
	log.Println(err)

	err = service.Error()
	log.Println(err)
}
