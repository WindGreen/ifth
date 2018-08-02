package main

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConfigDecode(t *testing.T) {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		t.Fatal(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", config)
}
