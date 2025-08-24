package config

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	AppPort      string `mapstructure:"APP_PORT"`
	AppEnv       string `mapstructure:"APP_ENV"`
	JwtSecretKey string `mapstructure:"JWT_SECRET"`
	JwtIssuer    string `mapstructure:"JWT_ISSUER"`
}

type PsqlDB struct {
    Host      string `mapstructure:"DATABASE_HOST"`
    Port      string `mapstructure:"DATABASE_PORT"`
    User      string `mapstructure:"DATABASE_USER"`
    Password  string `mapstructure:"DATABASE_PASSWORD"`
    DBName    string `mapstructure:"DATABASE_NAME"`
    SSLMode   string `mapstructure:"DATABASE_SSLMODE"`
    DBMaxOpen int    `mapstructure:"DATABASE_MAX_OPEN_CONNECTIONS"`
    DBMaxIdle int    `mapstructure:"DATABASE_MAX_IDLE_CONNECTIONS"`
}

type Config struct {
    AppPort   string `mapstructure:"APP_PORT"`
    AppEnv    string `mapstructure:"APP_ENV"`
    JwtSecret string `mapstructure:"JWT_SECRET"`
    JwtIssuer string `mapstructure:"JWT_ISSUER"`

    DBHost     string `mapstructure:"DATABASE_HOST"`
    DBPort     string `mapstructure:"DATABASE_PORT"`
    DBUser     string `mapstructure:"DATABASE_USER"`
    DBPassword string `mapstructure:"DATABASE_PASSWORD"`
    DBName     string `mapstructure:"DATABASE_NAME"`
    SSLMode    string `mapstructure:"DATABASE_SSLMODE"`
    DBMaxOpen  int    `mapstructure:"DATABASE_MAX_OPEN_CONNECTIONS"`
    DBMaxIdle  int    `mapstructure:"DATABASE_MAX_IDLE_CONNECTIONS"`
}


func LoadConfig(path string) (*Config, error) {
    viper.SetConfigFile(".env") // langsung tunjuk file
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Println("Warning: no .env file found, fallback to system ENV")
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}


