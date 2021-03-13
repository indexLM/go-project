package config

import (
	"go-project/utils"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	var MyServer *Server
	dir, err := os.Getwd()
	if err != nil {

	}
	err = utils.CofParse(dir+"/config.yaml", *MyServer)
	t.Log(MyServer)
}
