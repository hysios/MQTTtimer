package mqtttimer

import (
	"log"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tj/assert"
)

func mockClient() mqtt.Client {
	var opts = mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts.OnConnect = func(c mqtt.Client) {
		log.Printf("connected")
	}
	cli := mqtt.NewClient(opts)
	token := cli.Connect()
	token.Wait()
	return cli
}

func mockServer() *Server {
	var s = NewServe(mockClient())
	return s
}

func TestServer_Start(t *testing.T) {
	var s = NewServe(mockClient())
	go t.Run("serve", func(tt *testing.T) {
		s.Start()
		defer s.Stop()
	})
	time.Sleep(10 * time.Millisecond)
	cli := NewTimer(mockClient())
	err := cli.Sync()
	assert.NoError(t, err)
}
