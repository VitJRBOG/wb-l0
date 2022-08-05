package server

import (
	"fmt"
	"net/http"
	"os"
	"wb-l0/config"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Run(serverCfg *config.ServerConfig, dbConnection *gorm.DB) {
	rtr := mux.NewRouter()
	handling(rtr, dbConnection)

	http.Handle("/", rtr)
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverCfg.ServerPort), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable start the server: %v\n", err)
		os.Exit(1)
	}
}
