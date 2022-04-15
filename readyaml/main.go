package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//type _instance map[string]string
//type Ec2Instances struct {
//	_instance []_instance `yaml:"instances"`
//}

type Ec2Instances struct {
	Instances []struct {
		InstanceID string `yaml:"instanceId"`
		Service    string `yaml:"service"`
	} `yaml:"instances"`
}

func (ins *Ec2Instances) GetInstances(yamlPath string) *Ec2Instances {
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Fatalf("read yaml err, %v", err)
	}
	err = yaml.Unmarshal(yamlFile, ins)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return ins
}

func main() {
	var ins Ec2Instances

	instancesStruct := ins.GetInstances("instances.yaml")
	fmt.Printf("%v\n", instancesStruct.Instances)

}
