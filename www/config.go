package main

type Config struct {
	MongoDB struct {
		Host string `yaml:"host"`
	} `yaml:"mongodb"`
	WWW struct {
		Home string `yaml:"home"`
		Port int    `yaml:"port"`
	} `yaml:"www"`
	Url struct {
		Base      string `yaml:"base"`
		Length    int    `yaml:"length"`
		Unique    bool   `yaml:"unique`
		Algorithm int    `yaml:"algorithm"`
		Humanity  bool   `yaml:"humanity"`
	}
}
