package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func handling(rtr *mux.Router, dbConnection *gorm.DB) {
	rtr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := "Hello world"
		_, err := w.Write([]byte(text))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("GET", "POST")

	rtr.HandleFunc("/order/{order_uid:[a-zA-Z0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		orderUID := mux.Vars(r)["order_uid"]
		orderJSON, err := order(dbConnection, orderUID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(orderJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("GET", "POST")
}
