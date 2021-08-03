package mntp

import (
	"io"
	"log"
	"path"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	DefaultPrefix  = "MNTP"
	DefaultTimeout = 15 * time.Second
)

type Client struct {
	Prefix      string
	WaitTimeout time.Duration
	mqClient    mqtt.Client
}

type NtpPackage struct {
	T0, T1, T2, T3 int64
	Time           time.Time
}

func NewNTP(m mqtt.Client) *Client {
	return &Client{
		Prefix:      DefaultPrefix,
		WaitTimeout: DefaultTimeout,
		mqClient:    m,
	}
}

func (client *Client) Sync() error {
	var (
		sessid = UID()
		t      = time.Now()
		pkt    = NtpPackage{T0: t.UnixNano()}
		done   = make(chan bool)
	)

	client.mqClient.Subscribe(client.Topic("sessions", sessid), 0, func(c mqtt.Client, m mqtt.Message) {
		var t = time.Now()
		p := unpack(m.Payload())
		p.T3 = t.UnixNano()
		offset := ((p.T1 - p.T0) + (p.T2 - p.T3)) / 2
		t1 := p.T1 + offset
		nt := time.Unix(t1/1000000000, t1%10000000000)
		log.Printf("offset %v %s => %s", offset, p.Time, nt)

		c.Unsubscribe(client.Topic("sessions", sessid))
		done <- true
	})

	log.Printf("pack %s", pack(pkt))
	token := client.mqClient.Publish(client.Topic("synctime", sessid), 0, false, pack(pkt))
	token.Wait()

	select {
	case <-done:
		return nil
	case <-time.After(client.WaitTimeout):
		client.mqClient.Unsubscribe(client.Topic("sessions", sessid))
	}

	return io.EOF
}

func (client *Client) Topic(suffix ...string) string {
	ps := append([]string{client.Prefix}, suffix...)
	log.Printf("topic %s", path.Join(ps...))
	return path.Join(ps...)
}
