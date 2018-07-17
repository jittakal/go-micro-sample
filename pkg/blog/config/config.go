package config

import (
	"flag"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type (
	// Config holds application configuration
	Config struct {
		MongoDB MongoDB
		Port    int
	}

	// MongoDB related configuration within AppConfiguration
	MongoDB struct {
		App  App
		Auth Auth
	}

	// App related configuration within mongodb
	App struct {
		Database string
	}

	// Auth related configuration within mongodb
	Auth struct {
		Servers  string
		Database string
		Username string
		Password string
	}
)

// Config holds application level configuration
var config Config

func init() {
	// look for configuration
	viper.SetConfigName("config")      // name of config file (without extension)
	viper.AddConfigPath("/etc/blog/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.blog") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		log.Fatalf("unable to read application configuration, %v", err)
	}

	// using standard library "flag" package
	flag.Int("port", 50052, "server port")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// load configuration into structure
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}

// GetConfig return configuration
func GetConfig() Config {
	return config
}
