package main

import "wildberriesL0/pkg/server"

func main() {
	/*
		dbConn := postgres.NewClient()
		defer dbConn.Close()

		sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
		if err != nil {
			fmt.Println(err)
		}

		sub, _ := sc.Subscribe("orders", func(m *stan.Msg) {
			fmt.Printf("Message: %s\n", string(m.Data))
			var order models.Order
			err := json.Unmarshal(m.Data, &order)

			if err != nil {
				fmt.Errorf("New order not added: %s", err.Error())
				return
			}

			if order.UID == "" {
				fmt.Println("Error: empty user ID")
				return
			}

			_, err = dbConn.Exec("insert into orders (uid,order_fields) values ($1,$2);", order.UID, m.Data)
			if err != nil {
				fmt.Println("Insert into DB failed", err.Error())
			}
			fmt.Println("Success")
		})

		defer sub.Close()
		defer sc.Close()*/
	server.CreateServer()
	/*err := conn.QueryRow("SELECT * FROM orders").Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name)*/
}
