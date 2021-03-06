package mqtttimer

import (
	"path"
	"time"

	"github.com/hysios/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Server struct {
	Prefix   string
	UseUTC   bool
	mqClient mqtt.Client
	done     chan bool
}

func NewServe(m mqtt.Client) *Server {
	return &Server{Prefix: DefaultPrefix, mqClient: m, done: make(chan bool)}
}

func (s *Server) now() time.Time {
	if s.UseUTC {
		return utc()
	} else {
		return now()
	}
}

func (s *Server) Start() error {
	s.mqClient.Subscribe(s.Topic("synctime/+"), 2, func(c mqtt.Client, m mqtt.Message) {
		var (
			t      = s.now()
			sessid = path.Base(m.Topic())
			p      = unpack(m.Payload())
		)
		defer m.Ack()
		p.T1 = t.UnixNano()
		p.Time = s.now()
		time.Sleep(2 * time.Millisecond)
		p.T2 = utc().UnixNano()
		log.Debugf("time %s", t)
		s.mqClient.Publish(s.Topic("sessions", sessid), 2, false, pack(p))
	})

	select {
	case <-s.done:
		return nil
	}
}

func (s *Server) Stop() error {
	s.done <- true
	return nil
}

func (s *Server) Topic(suffix ...string) string {
	ps := append([]string{s.Prefix}, suffix...)
	log.Debugf("topic %s", path.Join(ps...))

	return path.Join(ps...)
}
