package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Cfg struct {
	Address string
	Port    int
}

func Parse() (result *Cfg, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetEnvPrefix("WEBSERVICE")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	err = viper.Unmarshal(&result)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into worker, %v", err)
	}

	// err = result.Validate()
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid configuration provided %v", err)
	// }

	return
}
