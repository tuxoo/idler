package config

import (
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	defaultHttpPort           = "8000"
	defaultHttpRWTimeout      = 10 * time.Second
	defaultMaxHeaderMegabytes = 1
	defaultTokenTTL           = 30 * time.Minute
)

type (
	Config struct {
		HTTP     HTTPConfig
		Auth     AuthConfig
		Postgres PostgresConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}

	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		TokenTTL   time.Duration
		SigningKey string
	}

	PostgresConfig struct {
		Host     string
		Port     string
		DB       string
		User     string
		Password string
		SSLMode  string
	}
)

func Init(path string) (*Config, error) {
	preDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func preDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
	viper.SetDefault("auth.tokenTTL", defaultTokenTTL)
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // config file name

	return viper.ReadInConfig()
}

func parseEnv() error {
	if err := parseLineEnv("jwt", "signing_key"); err != nil {
		return err
	}

	if err := parseLineEnv("http", "host"); err != nil {
		return err
	}

	if err := parsePostgresEnv(); err != nil {
		return err
	}

	return parseLineEnv("password", "salt")
}

func parseLineEnv(prefix, name string) error {
	viper.SetEnvPrefix(prefix)
	return viper.BindEnv(name)
}

func parsePostgresEnv() error {
	viper.SetEnvPrefix("postgres")

	if err := viper.BindEnv("host"); err != nil {
		return err
	}

	if err := viper.BindEnv("port"); err != nil {
		return err
	}

	if err := viper.BindEnv("db"); err != nil {
		return err
	}

	if err := viper.BindEnv("user"); err != nil {
		return err
	}

	if err := viper.BindEnv("password"); err != nil {
		return err
	}

	return viper.BindEnv("sslmode")
}

// Unmarshal config.yml by keys
func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Auth.PasswordSalt = viper.GetString("salt")
	cfg.Auth.JWT.SigningKey = viper.GetString("signing_key")

	cfg.HTTP.Host = viper.GetString("host")

	cfg.Postgres.Host = viper.GetString("host")
	cfg.Postgres.Port = viper.GetString("port")
	cfg.Postgres.DB = viper.GetString("db")
	cfg.Postgres.User = viper.GetString("user")
	cfg.Postgres.Password = viper.GetString("password")
}
