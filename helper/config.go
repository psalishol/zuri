package helper

import (
	"github.com/spf13/viper"
)



type config struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	Dsn string `mapstructure:"DSN"`
}

func ReadConfig (path string) (conf config, err  error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")  
	viper.AddConfigPath("env")

	viper.AutomaticEnv()

	if viper.ReadInConfig(); err != nil {
		return;
	}

	if viper.Unmarshal(&conf); err != nil {
		return;
	}

	return;
}