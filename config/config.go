package config

import (
	"github.com/spf13/viper"
)

func ParseConfig() *viper.Viper { //возможность конфигурирования оснывных параметров приложения (Адрес API для проверки, Порт на которм будет подниматься сервер, Уровень логирования)
	v := viper.New()
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	v.SetConfigFile("config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil
	}
	return v
}
