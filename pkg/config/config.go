package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type DatasourceType string

const (
	Mqtt DatasourceType = "mqtt"
)

type SensorRole string

const (
	Top             SensorRole = "top"
	Bottom                     = "bottom"
	Exhaust                    = "exhaust"
	Powerline                  = "powerline"
	Return                     = "return"
	RoomTemperature            = "room_t"
)

type SensorType string

const (
	Temperature SensorType = "temp"
)

type System string

const (
	Buffer System = "buffer"
	Boiler        = "boiler"
	Heater        = "heater"
)

type SensorConfig struct {
	Name          string         `yaml:"name"`
	Datasource    DatasourceType `yaml:"source"`
	Type          SensorType     `yaml:"type"`
	Address       string         `yaml:"address"`
	System        System         `yaml:"system"`
	Role          SensorRole     `yaml:"role"`
	MaxReadingAge string         `yaml:"max-reading-age"`
}

type ReportConfig struct {
	Type    DatasourceType `yaml:"type"`
	Address string         `yaml:"address"`
	System  System         `yaml:"system"`
	Role    SensorRole     `yaml:"role"`
}

type Config struct {
	Filepath string
	Sensors  []SensorConfig `yaml:"sensors"`
	Reports  []ReportConfig `yaml:"reports"`
}

func (c *Config) Load() error {
	f, err := os.Open(c.Filepath)
	if err != nil {
		return fmt.Errorf("config load failed with following reason: %w", err)
	}

	decoder := yaml.NewDecoder(f)
	decoder.Decode(c)
	defer f.Close()

	return nil
}
