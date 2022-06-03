package boot

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DefaultMysql *gorm.DB
	DefaultViper *viper.Viper
	Sugar        *zap.SugaredLogger
	DefaultLog   *zap.Logger
	DefaultConfig       *Config
)


func Run() error {
	if err := LoadConfig(); err != nil {
		return err
	}

	if err := InitLogger(); err != nil {
		return err
	}

	if err := InitMysql(); err != nil {
		return err
	}

	return nil

}
