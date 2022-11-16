package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver                string `mapstructure:"DB_DRIVER"`
	DBUrl                   string `mapstructure:"DATABASE_URL"`
	Host                    string `mapstructure:"HOST"`
	Port                    int    `mapstructure:"PORT"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	JwtSecretKey            string `mapstructure:"JWT_SECRET_KEY"`
	AccessTokenExpiredTime  int32  `mapstructure:"ACCESS_TOKEN_EXPIRED_TIME"`
	RefreshTokenExpiredTime int32  `mapstructure:"REFRESH_TOKEN_EXPIRED_TIME"`

	FBKey    string `mapstructure:"FB_KEY"`
	FBSecret string `mapstructure:"FB_SECRET"`
	GGKey    string `mapstructure:"GLE_KEY"`
	GGSecret string `mapstructure:"GLE_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
