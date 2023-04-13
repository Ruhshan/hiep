package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func readProperties()  {
	fmt.Println("Reading props")
	viper.AddConfigPath("backend/config")
	viper.AddConfigPath("config")
	viper.SetConfigName("dev")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Unable to read config")
	}
}


func Get(key string) string {
	if len(viper.AllKeys())==0{
		readProperties()
	}
	return viper.GetString(key)
}