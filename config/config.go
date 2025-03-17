package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//struct tags

var GlobalConfig Config

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

func LoadConfig() {
	//Named Return Parameters
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//viper.AutomaticEnv()
	err := viper.ReadInConfig()
	fmt.Println("read config")
	if err != nil {
		panic(fmt.Errorf("fatal errUtils config file: %w", err))
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		panic(err)
	}
}
