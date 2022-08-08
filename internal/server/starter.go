package server

import (
	"fmt"
	"net/http"
	"os"
	"wb-l0/config"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func Run(serverCfg *config.ServerConfig, redisClient *redis.Client) {
	rtr := mux.NewRouter()
	handling(rtr, redisClient)

	http.Handle("/", rtr)
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverCfg.ServerPort), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable start the server: %v\n", err)
		os.Exit(1)
	}
}
