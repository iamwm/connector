package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

func InitMongoConfig(configPath string) (*MongoConfig, error) {
	var mongoConfigInfo = &MongoConfig{}
	log.Printf("Mongo init task started")
	if configPath == "" {
		log.Printf("use default mongo config path")
		workPath, err := os.Getwd()
		if err != nil {
			return mongoConfigInfo, err
		}
		configPath = filepath.Join(workPath, "conf", "mongo.yaml")
	}
	if f, err := os.Open(configPath); err != nil {
		return mongoConfigInfo, err
	} else {
		err := yaml.NewDecoder(f).Decode(mongoConfigInfo)
		if err != nil {
			return mongoConfigInfo, err
		}
		return mongoConfigInfo, nil
	}
}

func InitRabbitmqConfig(configPath string) (*MqConfig, error) {
	var rabbitmqConfigInfo = &MqConfig{}
	log.Printf("Rabbitmq init task started")
	if configPath == "" {
		log.Printf("use default rabbitmq config path")
		workPath, err := os.Getwd()
		if err != nil {
			return rabbitmqConfigInfo, err
		}
		configPath = filepath.Join(workPath, "conf", "rabbitmq.yaml")
	}
	if f, err := os.Open(configPath); err != nil {
		return rabbitmqConfigInfo, err
	} else {
		err := yaml.NewDecoder(f).Decode(rabbitmqConfigInfo)
		if err != nil {
			return rabbitmqConfigInfo, err
		}
		return rabbitmqConfigInfo, nil
	}
}
