package mntp

import (
	"io"
	"path"
	"time"

	"github.com/hysios/log"

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
		t      = now()
		pkt    = NtpPackage{T0: t.UnixNano()}
		done   = make(chan bool)
	)

	client.mqClient.Subscribe(client.Topic("sessions", sessid), 0, func(c mqtt.Client, m mqtt.Message) {
		var t = now()
		p := unpack(m.Payload())
		p.T3 = t.UnixNano()
		offset := ((p.T1 - p.T0) + (p.T2 - p.T3)) / 2
		t1 := p.T0 + offset
		nt := time.Unix(t1/1000000000, t1%10000000000)
		log.Debugf("θ %d serve time %s", offset, p.Time)
		log.Debugf("diff %s %s => %s", nt.Sub(t), t, nt)

		SetSystemDate(nt)
		c.Unsubscribe(client.Topic("sessions", sessid))
		done <- true
	})

	log.Debugf("ts %d pack %s", t.UnixNano()/1000000, pack(pkt))
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
	log.Debugf("topic %s", path.Join(ps...))
	return path.Join(ps...)
}
