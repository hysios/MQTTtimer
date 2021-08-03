package main

import (
	"flag"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/hysios/mntp"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", "tcp://127.0.0.1:1883", "mqtt server broker addr")
}

func main() {
	flag.Parse()

	var (
		opts     = mqtt.NewClientOptions().AddBroker(addr)
		mqClient = mqtt.NewClient(opts)
	)
	mqClient.Connect()
	client := mntp.NewNTP(mqClient)
	var tick = time.NewTicker(10 * time.Second)
	for range tick.C {
		client.Sync()
	}
}
