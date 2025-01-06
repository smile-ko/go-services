package initialize

import (
	"go-services/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // Path to config
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// Read file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Unmarshal config
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(err)
	}
}
