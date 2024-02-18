package config

import "github.com/spf13/viper"

type Config struct {
	MQ       Rabbit
	Database Mongo
}

type Mongo struct {
	URI      string `mapstructure:"DB_URI"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_DATABASE"`
}
type Rabbit struct {
	Username string `mapstructure:"RABBIT_USERNAME"`
	Password string `mapstructure:"RABBIT_PASSWORD"`
	Host     string `mapstructure:"RABBIT_HOST"`
	Port     int    `mapstructure:"RABBIT_PORT"`
}

func NewConfig(folder, filename, envname string) (*Config, error) {
	cfg := new(Config)

	v := viper.New()

	v.AddConfigPath(folder)
	v.SetConfigName(filename)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	v.SetConfigFile(envname + ".env")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg.Database); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg.MQ); err != nil {
		return nil, err
	}

	return cfg, nil
}
