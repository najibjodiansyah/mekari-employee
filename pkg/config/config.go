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
		Database    string `mapstructure:"database"`
		Host        string `mapstructure:"host"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		MaxConn     int    `mapstructure:"max_conn"`      // max count of connections stored in the connection pool (the connection pool is maintained by go-pg)
		MaxIdleConn int    `mapstructure:"max_idle_conn"` // max count of connection in idle state in the connection pool, when the count of connection in idle state is more than this config, the excess connection will be closed
		Timeout     int    `mapstructure:"timeout"`       // timeout tolerance for dial attempt to database server (in seconds)
		PoolTimeout int    `mapstructure:"pool_timeout"`  // timeout tolerance for fetching a connection from the connection pool (in seconds)
		Prefix      string `mapstructure:"prefix"`
	} `mapstructure:"postgres"`
	GooglePubSubConfig GooglePubSubConfig `mapstructure:"google_pub_sub_config"`
	MPAdapter          MPAdapterConfig    `mapstructure:"mp_adapter"`
	PushOrderMpaCfg    PushOrderMpaCfg    `mapstructure:"push_order_mpa_cfg"`
	Warpath            struct {
		Host string `mapstructure:"host"`
	}
}

type MPAdapterConfig struct {
	APIKey         string `mapstructure:"api_key"`
	BaseURL        string `mapstructure:"base_url"`
	ChannelCode    string `mapstructure:"channel_code"`
	AwbFromChannel bool   `mapstructure:"awb_from_channel"`
}

type PushOrderMpaCfg struct {
	Worker int `mapstructure:"worker"`
}

type GooglePubSubConfig struct {
	MaxAckDeadlineSec       int64  `mapstructure:"max_ack_deadline_sec"`
	Type                    string `mapstructure:"type"`
	ProjectID               string `mapstructure:"project_id"`
	PrivateKeyID            string `mapstructure:"private_key_id"`
	PrivateKey              string `mapstructure:"private_key"`
	ClientEmail             string `mapstructure:"client_email"`
	ClientID                string `mapstructure:"client_id"`
	AuthURI                 string `mapstructure:"auth_uri"`
	TokenURI                string `mapstructure:"token_uri"`
	AuthProviderX509CertURL string `mapstructure:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `mapstructure:"client_x509_cert_url"`
	MaxOutstandingMessages  int    `mapstructure:"max_outstanding_messages"`
}

func init() {
	var err error

	viper.SetEnvPrefix("MKR")
	viper.AutomaticEnv()

	configName := "application-mekari-employee"

	if viper.IsSet("ENV") {
		Environment = viper.GetString("ENV")
	}

	if Environment != "production" {
		configName = configName + "." + Environment
	}

	log.Info("GOM_ENV: ", Environment)

	viper.SetConfigName(configName) // name of config file (without extension)
	viper.AddConfigPath(filepath.Join(GetAppBasePath(), "conf"))
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
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
