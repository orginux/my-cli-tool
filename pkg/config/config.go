package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Dest     string
	Emoticon map[string]string
}

func (c *Config) Load(cmd *cobra.Command) error {
	v := viper.New()
	v.BindPFlags(cmd.Flags())
	return nil
}

func Load() (*Config, error) {
	v := viper.New()
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = v.Unmarshal(c)
	return c, err
}
