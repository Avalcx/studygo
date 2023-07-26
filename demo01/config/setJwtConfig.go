package config

import (
	"strconv"

	"github.com/spf13/viper"
)

var (
	JwtSignKey string
	// 单位: 小时
	JwtExpiresTime int
)

func setJwtConfig() {
	// 签名key,默认值
	viper.SetDefault("JWT_SIGN_KEY", "abdas")
	// token有效期
	viper.SetDefault("JWT_EXP_TIME", "24")
	viper.AutomaticEnv()
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpiresTime, _ = strconv.Atoi(viper.GetString("JWT_EXP_TIME"))
}
