package main

import (
	"rosary/server"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Info("Starting Server")
	err := server.StartServer()
	if err != nil {
		logrus.Errorf("Server could not be started: %v\n", err)
	}

}
