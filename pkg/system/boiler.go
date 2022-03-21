package system

import (
	"fmt"
	"github.com/coff/go-sensor"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func main() {
	options := sensor.Options{Name: "exhaust", MaxReadingAge: 60 * time.Second}
	sensor1 := sensor.NewMqttSensor(&options)

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

	token := client.Subscribe("home/heating/boiler/exhaust/temp", 1, sensor1.Update)
	token.Wait()

	fmt.Print("Awaiting sensor1 initialization")

	for ok := true; ok; ok = sensor1.State == sensor.Inactive {
		time.Sleep(2 * time.Second)
		fmt.Print(".")
	}
	fmt.Println(" ok")

	for {
		reading, state, err := sensor1.Reading()

		if err == nil {
			fmt.Printf("Sensor %s reads %f, state %s\n", sensor1.Name(), reading, state)
		}

		time.Sleep(2 * time.Second)
	}

	client.Disconnect(250)
}
