package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	var MyServer *Server
	dir, err := os.Getwd()
	if err != nil {

	}
	err = CofParse(dir+"/config.yaml", *MyServer)
	t.Log(MyServer)
}
