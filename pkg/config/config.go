package config

import (
	"github.com/krau/favpics-helper/pkg/util"
	"github.com/spf13/viper"
)

type Config struct {
	Proxy struct {
		Enabled bool   `toml:"Enabled"`
		Addr    string `toml:"Addr"`
	} `toml:"proxy"`
	Sources struct {
		Pixiv struct {
			Enabled     bool   `toml:"Enabled"`
			RssURL      string `toml:"RssURL"`
			RefreshTime int    `toml:"RefreshTime"`
		} `toml:"Pixiv"`
	} `toml:"sources"`
	Storages struct {
		TelegramChannel struct {
			Enabled  bool   `toml:"Enabled"`
			UserName string `toml:"UserName"`
			ChatID   int64  `toml:"ChatID"`
		} `toml:"TelegramChannel"`
	} `toml:"storages"`
	Middlewares struct {
		TelegramBot struct {
			Enabled bool   `toml:"Enabled"`
			Token   string `toml:"Token"`
		} `toml:"TelegramBot"`
	} `toml:"middlewares"`
}

var Conf *Config = new(Config)

func init() {
	util.Log.Debug("init config")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("../../")
	viper.AddConfigPath(".") // optionally look for config in the working directory
	viper.SetConfigType("toml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		util.Log.Fatal("read config failed: %v", err)
	}
	viper.Unmarshal(&Conf)
	util.Log.Debug("init config success")
}

/*
func ReadConfig() (Config, error) {
	util.Log.Debug("read config")
	var c Config
	configFile := utils.GetCurrentAbPath() + "/test.toml"
	util.Log.Debugf("config file path: %s", configFile)
	_, err := toml.DecodeFile(configFile, &c)
	if err != nil {
		return Config{}, err
	}
	util.Log.Debug("read config success")
	return c, err
}
*/
