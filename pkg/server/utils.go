package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"wildberriesL0/pkg/models"
)

func FindOrderInCache(id string, cache map[string]models.Order) models.Order {
	value, ok := cache[id]
	if ok {
		return value
	}
	return models.Order{}
}

func GetCache(db *sql.DB, cache map[string]models.Order) {
	rows, err := db.Query("select order_fields from orders where order_fields is not null")
	if err != nil {
		log.Fatalln("Error from query: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		row := []byte{}
		order := models.Order{}
		err = rows.Scan(&row)
		if err != nil {
			log.Fatalln("Error with scan: ", err)
		}

		err = json.Unmarshal(row, &order)
		if err != nil {
			log.Fatalln("Error with unmarshal:", err)
		}
		cache[order.UID] = order
	}
}
