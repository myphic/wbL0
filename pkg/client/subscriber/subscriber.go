package subscriber

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"wildberriesL0/pkg/models"
)

func Subscribe(db *sql.DB, cache map[string]models.Order) {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("subscriber://localhost:4222"))
	if err != nil {
		log.Println(err)
	}

	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		fmt.Printf("Message: %s\n", string(m.Data))
		var order models.Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Fatalf("New order not added: %s", err.Error())
			return
		}

		if order.UID == "" {
			log.Println("Error: empty user ID")
			return
		}

		_, err = db.Exec("INSERT INTO orders (uid,order_fields) VALUES ($1,$2);", order.UID, m.Data)
		if err != nil {
			log.Println("Insert into DB failed", err.Error())
		}
		cache[order.UID] = order
		log.Println("Success")
	})
	if err != nil {
		log.Fatalf("error: %s", err)
	}

}
