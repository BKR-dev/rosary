package main

import (
	"rosary/server"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Info("Starting Server")
	server.StartServer()
}
