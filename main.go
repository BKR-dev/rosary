package main

import (
	"fmt"
	"rosary/prayers"
	"rosary/server"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Info("Starting Server")

	rosary, _ := prayers.DailyRosary()
	fmt.Println(rosary)

	err := server.StartServer()
	if err != nil {
		logrus.Errorf("Server could not be started: %v\n", err)
	}

}
