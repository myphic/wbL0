package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("json/model.json")
	if err != nil {
		log.Fatalln("Cant read the file", err, file)
	}
	sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		fmt.Println(err)
	}

	err = sc.Publish("orders", file)
	if err != nil {
		fmt.Println(err)
	}

	defer sc.Close()

}
