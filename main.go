package main

import (
	"fmt"
	"github.com/raj47i/go-utils/config"
)

// Configuration holds the app configuration, struct makes it easier to load & save it as json
type Configuration struct {
	LogLevel string `json:"log_level" env:"LOG_LEVEL" default:"info"`
	Debug    bool   `json:"debug" env:"DEBUG" default:"true"`
}

func main() {
	var cfg Configuration
	fmt.Println(config.LoadFromENV(&cfg))

	fmt.Println(cfg)
}
