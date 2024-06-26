package internal

import (
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"gopkg.in/yaml.v2"
	"os"
)

// Config represents the configuration object for the application.
type Config struct {
	Seeding *Seeding `yaml:"seeding"`
}

// Seeding represents the configuration with predefined seeds.
type Seeding struct {
	Dummies      *Dummies           `yaml:"dummies"`
	Instruments  []*inst.Instrument `yaml:"instruments"`
	Difficulties []*diff.Difficulty `yaml:"difficulties"`
	Sources      []*src.Source      `yaml:"sources"`
	Endpoints    []*end.Endpoint    `yaml:"endpoints"`
}

// Dummies represents the configuration for generating dummies.
type Dummies struct {
	Artists *Artists `yaml:"artists"`
	Tracks  *Tracks  `yaml:"tracks"`
	Tabs    *Tabs    `yaml:"tabs"`
}

// Artists represents a struct for seeding artists.
type Artists struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// Tracks represents a struct for seeding tracks.
type Tracks struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// Tabs represents a struct for seeing tabs.
type Tabs struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// NewConfig reads a file from the given path and unmarshalls its contents into a Config struct.
func NewConfig(path string) (*Config, error) {
	var config Config

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
