package main

import (
	"flag"
	"os"
	"time"

	timer "github.com/hysios/MQTTtimer"
	"github.com/hysios/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	s := timer.NewServe(mqClient)
	log.Infof("startup mntp server connect %s", addr)
	s.Start()
}
