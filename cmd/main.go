package main

import (
	"wildberriesL0/pkg/client/postgres"
	"wildberriesL0/pkg/client/subscriber"
	"wildberriesL0/pkg/models"
	"wildberriesL0/pkg/server"
)

var cache map[string]models.Order

func main() {
	cache := map[string]models.Order{}
	dbConn := postgres.NewClient()
	defer dbConn.Close()
	server.GetCache(dbConn, cache)
	subscriber.Subscribe(dbConn, cache)

	CreateServer()
}
