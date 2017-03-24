package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/hyperpilotio/qos-data-store/api"
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

	s := api.NewServer(viper)
	return s.StartServer()
}

func main() {
	configPath := flag.String("config", "", "The file path to a config file")
	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Git Commit:", GitCommit)
		fmt.Println("Version:", Version)
		if VersionPrerelease != "" {
			fmt.Println("Version PreRelease:", VersionPrerelease)
		}
		return
	}

	err := Run(*configPath)
	if err != nil {
		glog.Errorln(err)
	}
}
