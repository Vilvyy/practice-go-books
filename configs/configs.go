package configs

import "github.com/spf13/viper"

func LoadEnv(envFile string) error {
	viper.SetConfigFile(envFile)
	return viper.ReadInConfig()
}
