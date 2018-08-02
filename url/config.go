package main

type Config struct {
	MongoDB struct {
		Host string `yaml:"host"`
	} `yaml:"mongodb"`
	WWW struct {
		Home string `yaml:"home"`
	}
	Url struct {
		Port int    `yaml:"port"`
		Base string `yaml:"base"`
	}
}
