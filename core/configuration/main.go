package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

type ADOConfiguration struct {
	PAT string `mapstructure:"PAT"`
}

type Configuration struct {
	ADO ADOConfiguration `mapstructure:"ADO"`
}

func LoadConfiguration() (configuration *Configuration, err error) {
	viper.SetEnvPrefix("RANGER")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.BindEnv("ado.pat")

	err = viper.Unmarshal(&configuration)

	return
}
