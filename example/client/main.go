package main

import (
	"flag"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/hysios/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	timer "github.com/hysios/MQTTtimer"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", "", "mqtt server broker addr")
}

func main() {
	flag.Parse()

	if len(addr) == 0 {
		addr = os.Getenv("SERVER_IP")
	}

	var (
		opts     = mqtt.NewClientOptions().AddBroker(addr)
		mqClient = mqtt.NewClient(opts)
	)

	opts.SetOnConnectHandler(func(c mqtt.Client) {
		log.Infof("connected")
	})

	if token := mqClient.Connect(); token.Wait() && token.Error() != nil {
		time.Sleep(5 * time.Second)
		panic(token.Error())
	}

	log.Infof("connect %s", addr)
	client := timer.NewTimer(mqClient)
	client.Sync()

	var tick = time.NewTicker(10 * time.Second)
	for range tick.C {
		client.Sync()
	}
}
