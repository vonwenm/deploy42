package config

import (
  "io/ioutil"
  "gopkg.in/yaml.v1"
)

type Configuration struct {
  Service map[string]interface{}
  Commands map[string][]map[string]interface{}
}

func unmarshal() Configuration {
  deploy := Configuration {}
  data, _ := ioutil.ReadFile("deploy.yml")

  yaml.Unmarshal(data, &deploy)
  return deploy
}

var config Configuration = unmarshal()

func NewCoreService() CoreService {
  service := config.Service
  return CoreService { service["port"].(int), service["bind"].(string) }
}

func NewCommands() []Command {

}
