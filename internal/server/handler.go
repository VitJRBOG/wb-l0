package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func handling(rtr *mux.Router, redisClient *redis.Client) {
	rtr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := "Hello world"
		_, err := w.Write([]byte(text))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("GET", "POST")

	rtr.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		t, err := getTemplate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.ExecuteTemplate(w, "order", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}).Methods("GET", "POST")

	rtr.HandleFunc("/order/get", func(w http.ResponseWriter, r *http.Request) {
		orderUID := r.FormValue("order_uid")
		orderJSON, err := order(redisClient, orderUID)
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

func getTemplate() (*template.Template, error) {
	rootPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	localPath := "internal/server/ui/html"
	path := filepath.FromSlash(fmt.Sprintf("%s/%s", rootPath, localPath))

	t, err := template.ParseFiles(fmt.Sprintf("%s/%s", path, "order.html"))
	if err != nil {
		return nil, err
	}

	return t, nil
}
