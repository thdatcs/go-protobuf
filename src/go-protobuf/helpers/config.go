package helpers

import (
	"go-protobuf/configs"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// InitConfig initializes config
func InitConfig(configFile string) (*configs.Config, error) {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config configs.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// GetString returns the value associated with the key as a string
func GetString(key string) string {
	value, err := os.LookupEnv(key)
	if !err {
		value = viper.GetString(key)
	}
	return value
}

// GetBool returns the value associated with the key as a bool
func GetBool(key string) (bool, error) {
	value := GetString(key)
	return strconv.ParseBool(value)
}

// GetInt returns the value associated with the key as an int
func GetInt(key string) (int, error) {
	value := GetString(key)
	return strconv.Atoi(value)
}
