package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DB struct {
	Uri  string `yaml:"uri"`
	Name string `yaml:"name"`
}

type Servers struct {
	AppServer string `yaml:"appServer"`
	DB        DB     `yaml:"db"`
}

type Configuration struct {
	Servers Servers `yaml:"servers"`
}

var fileName string = "./conf.yaml"
var conf Configuration

func Load() {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Printf("Configuration : %+v", conf)
}

func Get() Configuration {
	return conf
}
