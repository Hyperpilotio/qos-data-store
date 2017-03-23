package main

import (
	"errors"
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

// Run start the web server
func Run(fileConfig string) error {
	viper := viper.New()
	viper.SetConfigType("json")

	if fileConfig == "" {
		return errors.New("Config file path is required")
	}

	viper.SetConfigFile(fileConfig)
	viper.SetDefault("port", "7781")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("Unable to read config: " + err.Error())
	}

	server := NewServer(viper)
	return server.StartServer()
}

func main() {
	configPath := flag.String("config", "", "The file path to a config file")
	flag.Parse()

	err := Run(*configPath)
	glog.Errorln(err)
}
