# MQTTTimer is based mqtt protocol sync timer.

You can used ntp sync time protocol in IoT without ntp server.
used mqtt protocol sync time is tcp connnections, and based mqtt protocol, so don't need a ntp server.
and direct embed mqtt server side.

# exmaples

### Import 
```go
import timer "github.com/hysios/MQTTtimer"
```

### Client Mode

```go
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

	log.Printf("connect %s", addr)
	client := timer.NewTimer(mqClient)
	client.Sync()
```

### Server Mode

```go

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
	log.Printf("startup mntp server connect %s", addr)
	s.Start()
```
