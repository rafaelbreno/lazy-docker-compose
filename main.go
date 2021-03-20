package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

var dockerCompose map[interface{}]interface{}

// Docker Compose's main file
type DockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

// Docker Compose's services
type Service struct {
	Build         Build    `yaml:"build"`          // Service's build data
	ContainerName string   `yaml:"container_name"` // Service's container name
	DependsOn     []string `yaml:"depends_on"`     // Which services this service's build depends on
	Image         string   `yaml:"image"`          // Service's image's name/ID
	Ports         []string `yaml:"ports"`          // Service's ports
	Restart       string   `yaml:"restart"`        // When this service should restart
	Volumes       []string `yaml:"volumes"`        // Where this service will store it's volume
	Environment   []string `yaml:"environment"`    // Service's environment variables
}

// Service's build data
type Build struct {
	Context    string `yaml:"context"`    // Build context
	Dockerfile string `yaml:"dockerfile"` // Dockerfile path
}

func main() {
	file, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
		panic(err)
	}

	var e DockerCompose

	err = yaml.Unmarshal(file, &e)

	if err != nil {
		panic(err)
	}

	fmt.Println(e)
}
