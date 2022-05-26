package server

import (
	"wildberriesL0/pkg/client/postgres"
	"wildberriesL0/pkg/client/subscriber"
	"wildberriesL0/pkg/models"
)

var cache map[string]models.Order

func Register() {
	cache = make(map[string]models.Order)
	dbConn := postgres.NewClient()
	defer dbConn.Close()
	GetCache(dbConn)
	subscriber.Subscribe(dbConn, cache)
	CreateServer()
}
