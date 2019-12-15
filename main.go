package main

import (
	"fmt"
	"time"

	"github.com/raj47i/go-utils/daemon"
)

// Configuration holds the app configuration, struct makes it easier to load & save it as json
type Configuration struct {
	LogLevel string `json:"log_level" env:"LOG_LEVEL" default:"info"`
	Debug    bool   `json:"debug" env:"DEBUG" default:"true"`
}

func main() {

	s, err := daemon.Run("go-utils")
	if err == nil {
		fmt.Println("App started successfully.")
		fmt.Println(err)
		fmt.Println("Waiting 30 seconds")
		time.Sleep(30 * time.Second)
		daemon.Exit(s)
	} else if err == daemon.ErrAlreadyRunning {
		fmt.Println("Another Instance already running")
	} else {
		fmt.Println(err)
	}

}
