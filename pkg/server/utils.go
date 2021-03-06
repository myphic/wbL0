package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"wildberriesL0/pkg/models"
)

func FindOrderInCache(id string) (models.Order, error) {
	value, ok := cache[id]
	if ok {
		return value, nil
	}
	return models.Order{}, errors.New("Order not found")
}

func GetCache(db *sql.DB) {
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
			log.Fatalln("Error scan: ", err)
		}

		err = json.Unmarshal(row, &order)
		if err != nil {
			log.Fatalln("Error unmarshal:", err)
		}
		cache[order.UID] = order
	}
}
