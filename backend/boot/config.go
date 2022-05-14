package boot

import "github.com/spf13/viper"

func LoadConfig() error {
	config := "config.yaml"
	VIPER := viper.New()
	VIPER.SetConfigFile(config)
	VIPER.SetConfigType("yaml")
	err := VIPER.ReadInConfig()
	if err != nil {
		return err
	}
	if err := VIPER.Unmarshal(&CONFIG); err != nil {
		return err
	}
	return nil
}
