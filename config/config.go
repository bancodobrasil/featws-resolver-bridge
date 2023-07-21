package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config type contains various fields with corresponding mapstructure tags used for configuration
// in a Go application.
//
// Property:
//   - {string} Port: The port number on which the server will listen for incoming requests.
//   - {string} ResolversFile: is a property that specifies the file path for the resolvers used by the feature web service resolver bridge. Resolvers are functions that map GraphQL queries to their corresponding data sources. The ResolversFile property is used to configure the location of the file containing these resolver functions.
//   - {int64} ReadyTimeout: is a property of type int64 that represents the maximum amount of time in seconds that the resolver bridge will wait for all resolvers to become ready before starting to serve requests. If any resolver fails to become ready within this time, the resolver bridge will exit with an error.
//   - {string} ExternalHost: is a string that represents the external host name or IP address that the resolver bridge will use to communicate with external services. This property is typically used to configure the resolver bridge to communicate with services that are hosted on a different machine or network.
//   - {string} AuthAPIKey: is a string that represents the API key used for authentication in the resolver bridge.
type Config struct {
	Port          string `mapstructure:"PORT"`
	ResolversFile string `mapstructure:"FEATWS_RESOLVER_BRIDGE_RESOLVERS_FILE"`
	ReadyTimeout  int64  `mapstructure:"FEATWS_RESOLVER_BRIDGE_READY_TIMEOUT"`
	ExternalHost  string `mapstructure:"EXTERNAL_HOST"`

	AuthAPIKey string `mapstructure:"FEATWS_RESOLVER_BRIDGE_API_KEY"`
}

var config = &Config{}

// LoadConfig loads configuration settings from a file and sets default values if they are not
// present.
func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetDefault("PORT", "9000")
	viper.SetDefault("FEATWS_RESOLVER_BRIDGE_RESOLVERS_FILE", "./resolvers.json")
	viper.SetDefault("FEATWS_RESOLVER_BRIDGE_READY_TIMEOUT", 2)
	viper.SetDefault("EXTERNAL_HOST", "localhost:9000")
	viper.SetDefault("FEATWS_RESOLVER_BRIDGE_API_KEY", "")

	err = viper.ReadInConfig()
	if err != nil {
		if err2, ok := err.(*os.PathError); !ok {
			err = err2
			log.Errorf("Error on Load Config: %v", err)
			return
		}
	}

	err = viper.Unmarshal(config)

	return
}

// GetConfig returns a pointer to a Config object.
func GetConfig() *Config {
	return config
}
