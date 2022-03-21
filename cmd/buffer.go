package main

import (
	"fmt"
	"github.com/coff/go-hellfire/pkg/config"
	"github.com/coff/go-hellfire/pkg/system"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	cfg := config.Config{Filepath: "hellfire.yaml"}
	err := cfg.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var broker = "10.10.0.112"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	//opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetUsername("coff")
	opts.SetPassword("karroyo7")
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	bfr := system.Buffer{}
	bfr.Client = client
	err = bfr.Bootstrap(&cfg)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
