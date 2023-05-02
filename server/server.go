package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func StartServer() error {
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		logrus.Errorln(err)
		return err
	}
	return nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

}
