package system

import (
	"fmt"
	"github.com/coff/go-hellfire/pkg/config"
	"github.com/coff/go-sensor"
)

type Buffer struct {
	System
	SensorArray map[config.SensorRole]*sensor.ISensor
}

func (b *Buffer) Bootstrap(cfg config.Config) error {
	for _, sensorCfg := range cfg.Sensors {
		if sensorCfg.System != config.Buffer {
			continue
		}

		var newSensor *sensor.ISensor
		var err error
		switch sensorCfg.Datasource {
		case config.Mqtt:
			*newSensor, err = b.BootstrapMqttSensor(sensorCfg)
		}

		if err != nil {
			return fmt.Errorf("unable to register sensor %s; returned: %w", sensorCfg.Name, err)
		}

		b.SensorArray[sensorCfg.Role] = newSensor
	}

	return nil
}
