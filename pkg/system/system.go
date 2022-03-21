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
	Client mqtt.Client
}

func (s *System) BootstrapMqttSensor(sensorCfg *config.SensorConfig) (*sensor.MqttSensor, error) {

	duration, err := time.ParseDuration(sensorCfg.MaxReadingAge)

	if err != nil {
		return nil, fmt.Errorf("unable to parse max-reading-age: %s as time duration: %w", sensorCfg.MaxReadingAge, err)
	}

	opts := &sensor.Options{Name: sensorCfg.Name, MaxReadingAge: duration}
	fmt.Println(sensorCfg.Address)
	newSensor := sensor.NewMqttSensor(opts)

	s.Client.Subscribe(sensorCfg.Address, 1, newSensor.Update)

	return newSensor, nil
}
