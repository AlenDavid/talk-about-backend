package main

import (
	"github.com/alendavid/talk-about-backend/internal/broker"
	"github.com/alendavid/talk-about-backend/internal/database"
	"github.com/alendavid/talk-about-backend/internal/log"
	"github.com/alendavid/talk-about-backend/internal/network/http"
)

const (
	Host      = "database-service"
	Port      = 5432
	User      = "plants"
	Password  = "password"
	Name      = "plants"
	brokerUrl = "amqp://guest:guest@broker-service:5672/"
)

func main() {
	log.SetDefault()

	db, err := database.Run(Host, Port, User, Password, Name)
	if err != nil {
		panic(err)
	}

	broker, err := broker.Run(brokerUrl)
	if err != nil {
		panic(err)
	}

	defer broker.Close()

	network := http.New(db, *broker)

	// TODO: handle server shutdown better
	panic(network.Run())
}
