package mqtttimer

import (
	"io"
	"math"
	"path"
	"runtime"
	"time"

	"github.com/hysios/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	DefaultPrefix   = "MNTP"
	DefaultTimeout  = 15 * time.Second
	DefaultMaxRetry = 3
)

type Client struct {
	Prefix      string
	WaitTimeout time.Duration
	UseUTC      bool
	mqClient    mqtt.Client
}

type NtpPackage struct {
	T0, T1, T2, T3 int64
	Time           time.Time
}

func NewTimer(m mqtt.Client) *Client {
	return &Client{
		Prefix:      DefaultPrefix,
		WaitTimeout: DefaultTimeout,
		mqClient:    m,
	}
}

func (client *Client) now() time.Time {
	if client.UseUTC {
		return utc()
	} else {
		return now()
	}
}

func (client *Client) Sync() error {
	var (
		sessid = UID()
		t      = client.now()
		pkt    = NtpPackage{T0: t.UnixNano()}
		done   = make(chan bool)
	)

	token := client.mqClient.Subscribe(client.Topic("sessions", sessid), 2, func(c mqtt.Client, m mqtt.Message) {
		var (
			t = client.now()
			p = unpack(m.Payload())
		)
		p.T3 = t.UnixNano()
		offset := ((p.T1 - p.T0) + (p.T2 - p.T3)) / 2
		if runtime.GOOS == "windows" {
			t1 := p.T0 + offset
			nt := time.Unix(t1/1000000000, t1%10000000000)
			SetSystemDate(nt)
			log.Debugf("θ %d serve time %s", offset, p.Time)
			log.Debugf("diff %s %s => %s", nt.Sub(t), t, nt)
		} else {
			if math.Abs(float64(offset)) < 500000000 {
				Adjtime(offset)
				nt := client.now()
				log.Debugf("diff %s %s => %s", nt.Sub(t), t, nt)
				log.Infof("adjtime %d", offset)
			} else {
				t1 := p.T0 + offset
				nt := time.Unix(t1/1000000000, t1%10000000000)
				SetSystemDate(nt)
				log.Debugf("θ %d serve time %s", offset, p.Time)
				log.Debugf("diff %s %s => %s", nt.Sub(t), t, nt)
			}
		}
		c.Unsubscribe(client.Topic("sessions", sessid))
		done <- true
	})

	log.Infof("wait subscribe %v", token.WaitTimeout(30*time.Second))

	log.Debugf("ts %d pack %s", t.UnixNano()/1000000, pack(pkt))
	token = client.mqClient.Publish(client.Topic("synctime", sessid), 2, false, pack(pkt))
	token.WaitTimeout(30 * time.Second)
	log.Infof("wait publish %v", token.WaitTimeout(30*time.Second))

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
