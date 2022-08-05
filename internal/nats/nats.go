package nats

import (
	"fmt"
	"os"

	"wb-l0/config"

	"github.com/nats-io/stan.go"
)

func Connect(cfg *config.NatsConfig) stan.Conn {
	sc, err := stan.Connect(
		cfg.NatsClusterID,
		cfg.NatsClientID,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to Nats Streaming: %v\n", err)
		os.Exit(1)
	}

	return sc
}

func Subcribe(sc stan.Conn, msgChannel chan []byte) stan.Subscription {
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		msgChannel <- m.Data
	}, stan.StartWithLastReceived())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong on Nats Streaming: %v\n", err)
		os.Exit(1)
	}

	return sub
}
