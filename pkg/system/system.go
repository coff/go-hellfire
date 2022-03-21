package system

import (
	"fmt"
	"github.com/coff/go-hellfire/pkg/config"
	"github.com/coff/go-sensor"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type ISystem interface {
	Bootstrap(config config.Config)
}

type System struct {
	MqttClient mqtt.Client
}

func (s *System) BootstrapMqttClient(clientCfg *config.MqttClientConfig) {
	fmt.Println(clientCfg)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", clientCfg.Host, clientCfg.Port))
	opts.SetClientID(clientCfg.ClientId)
	//opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetUsername(clientCfg.Username)
	opts.SetPassword(clientCfg.Password)
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}
	s.MqttClient = mqtt.NewClient(opts)
	if token := s.MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Errorf("system: %w", token.Error()))
	}
}

func (s *System) BootstrapMqttSensor(sensorCfg *config.SensorConfig) (*sensor.MqttSensor, error) {

	duration, err := time.ParseDuration(sensorCfg.MaxReadingAge)

	if err != nil {
		return nil, fmt.Errorf("unable to parse max-reading-age: %s as time duration: %w", sensorCfg.MaxReadingAge, err)
	}

	opts := &sensor.Options{Name: sensorCfg.Name, MaxReadingAge: duration}
	fmt.Println(sensorCfg.Address)
	newSensor := sensor.NewMqttSensor(opts)

	s.MqttClient.Subscribe(sensorCfg.Address, 1, newSensor.Update)

	return newSensor, nil
}

func (s *System) Close() {
	if s.MqttClient.IsConnected() {
		s.MqttClient.Disconnect(2000)
	}
}
