package utils

import (
	"log"

	"gopkg.in/yaml.v2"
)

type YAML []byte

func (file YAML) Unmarshal(cfg interface{}) (interface{}, error) {
	if err := yaml.Unmarshal(file, cfg); err != nil {
		log.Printf(err.Error())
		return cfg, err
	}
	return cfg, nil
}