package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		viper.SetConfigFile("../.env")
		err = viper.ReadInConfig()
		if err != nil {
			return config, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func GetWeatherApiKey() string {
	config, _ := LoadConfig()
	return config.WeatherApiKey
}
