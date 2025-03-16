package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"log"
)

//struct tags

var globalConfig Config

type Config struct {
	Datasource Datasource `yaml:"datasource"`
	Server     Server     `yaml:"server"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Datasource struct {
	DbType   string `yaml:"dbType"`
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

func LoadConfig() (config Config, err error) {
	//Named Return Parameters
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//viper.AutomaticEnv()
	err = viper.ReadInConfig()

	fmt.Println("read config")

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("Config: %+v", spew.Sdump(config))

	return
}
