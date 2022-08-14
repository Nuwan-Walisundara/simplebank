package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver       string `mapstructure:"driverName"`
	Data_Source    string `mapstructure:"dataSourceName"`
	Server_Address string `mapstructure:"address"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
