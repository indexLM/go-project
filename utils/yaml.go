package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//读取yaml文件到对应类
func CofParse(file string, in interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, in)
	if err != nil {
		return err
	}
	return nil
}
