package boot

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Local   Local   `mapstructure:"local" json:"local" yaml:"local"`
}

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`       // 验证码长度
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`    // 验证码宽度
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"` // 验证码高度
}

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 端口值
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"` // 存放casbin模型的相对路径
}

func LoadConfig() error {
	config := "config.yaml"
	VIPER := viper.New()
	VIPER.SetConfigFile(config)
	VIPER.SetConfigType("yaml")
	err := VIPER.ReadInConfig()
	if err != nil {
		return err
	}
	if err := VIPER.Unmarshal(&DefaultConfig); err != nil {
		return err
	}
	return nil
}
