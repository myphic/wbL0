package main

import (
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("json/model.json")
	if err != nil {
		log.Fatalln("Cant read the file", err, file)
	}
	sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("subscriber://localhost:4222"))
	if err != nil {
		log.Printf("Stan connection error: %s", err)
	}

	err = sc.Publish("orders", file)

	if err != nil {
		log.Println(err)
	}
	log.Println("Stan published file into queue")
	defer sc.Close()

}
