package config

import (
	"fmt"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config config
var Environment string = "local"

type config struct {
	PgCfg struct {
		Database string `mapstructure:"database"`
		Host     string `mapstructure:"host"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Port     string `mapstructure:"port"`
	} `mapstructure:"postgres"`
}

func init() {
	var err error

	viper.SetEnvPrefix("MKR")
	viper.AutomaticEnv()

	configName := "app"

	if viper.IsSet("ENV") {
		Environment = viper.GetString("ENV")
	}

	if Environment != "production" && Environment != "dev" || Environment != "test" {
		configName = configName + "." + Environment
	}

	log.Info("MKR_ENV: ", Environment)

	viper.SetConfigName(configName) // name of config file (without extension)
	viper.AddConfigPath(filepath.Join(GetAppBasePath(), "conf"))
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	//Unmarshal application yml to config
	err = viper.Unmarshal(&Config)

	if err != nil {
		log.Errorf("unable to decode into struct, %v", err)
	}
}

func GetAppBasePath() string {
	basePath, _ := filepath.Abs(".")
	for filepath.Base(basePath) != "mekari-employee" {
		basePath = filepath.Dir(basePath)
	}
	return basePath
}
