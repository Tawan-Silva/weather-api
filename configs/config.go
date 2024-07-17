package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	GeoApiFyApiKey string `mapstructure:"GEOAPIFY_API_KEY"`
	WeatherApiKey  string `mapstructure:"WEATHER_API_KEY"`
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

func GetGeoApiKey() string {
	config, _ := LoadConfig()
	return config.GeoApiFyApiKey
}

func GetWeatherApiKey() string {
	config, _ := LoadConfig()
	return config.WeatherApiKey
}
