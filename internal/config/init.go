package config

import (
	"os"

	"github.com/Velocyes/mini-go-project/internal/consts"
	"github.com/Velocyes/mini-go-project/internal/model"
	"gopkg.in/yaml.v2"
)

func InitConfig() (config *model.Config, err error) {
	f, err := os.Open(consts.ConfigFilepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return
}
