package config

import (
	"github.com/eugene-krivtsov/idler/pkg/db/mongo"
	"github.com/eugene-krivtsov/idler/pkg/db/postgres"
	"github.com/eugene-krivtsov/idler/pkg/db/redis"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	defaultHttpPort           = "8080"
	defaultHttpRWTimeout      = 10 * time.Second
	defaultMaxHeaderMegabytes = 1
	defaultTokenTTL           = 30 * time.Minute
)

type (
	Config struct {
		HTTP     HTTPConfig
		Auth     AuthConfig
		Postgres postgres.Config
		Redis    redis.Config
		Mongo    mongo.Config
		WS       WSConfig
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

	WSConfig struct {
		Port            string
		ReadBufferSize  int
		WriteBufferSize int
	}
)

func Init(path string) (*Config, error) {
	viper.AutomaticEnv()
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

	if err := parseHttpEnv(); err != nil {
		return err
	}

	if err := parsePostgresEnv(); err != nil {
		return err
	}

	if err := parseRedisEnv(); err != nil {
		return err
	}

	if err := parseMongoEnv(); err != nil {
		return err
	}

	if err := parseLineEnv("password", "salt"); err != nil {
		return err
	}

	if err := parseLineEnv("websocket.port", "WEBSOCKET_PORT"); err != nil {
		return err
	}

	return nil
}

func parseMongoEnv() error {
	if err := viper.BindEnv("mongo.host", "MONGO_HOST"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.port", "MONGO_PORT"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.db", "MONGO_DB"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.user", "MONGO_INITDB_ROOT_USERNAME"); err != nil {
		return err
	}

	return viper.BindEnv("mongo.password", "MONGO_INITDB_ROOT_PASSWORD")
}

func parseLineEnv(prefix, name string) error {
	viper.SetEnvPrefix(prefix)
	return viper.BindEnv(name)
}

func parseHttpEnv() error {
	if err := viper.BindEnv("http.host", "HTTP_HOST"); err != nil {
		return err
	}

	return viper.BindEnv("http.port", "HTTP_PORT")
}

func parseRedisEnv() error {
	if err := viper.BindEnv("redis.host", "REDIS_HOST"); err != nil {
		return err
	}

	if err := viper.BindEnv("redis.port", "REDIS_PORT"); err != nil {
		return err
	}

	return viper.BindEnv("redis.password", "REDIS_PASSWORD")
}

func parsePostgresEnv() error {

	if err := viper.BindEnv("postgres.host", "POSTGRES_HOST"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.port", "POSTGRES_PORT"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.db", "POSTGRES_DB"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.user", "POSTGRES_USER"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.password", "POSTGRES_PASSWORD"); err != nil {
		return err
	}

	return viper.BindEnv("postgres.sslmode", "POSTGRES_SSLMODE")
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

	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("websocket", &cfg.WS); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Auth.PasswordSalt = viper.GetString("salt")
	cfg.Auth.JWT.SigningKey = viper.GetString("signing_key")

	cfg.HTTP.Host = viper.GetString("http.host")
	cfg.HTTP.Port = viper.GetString("http.port")

	cfg.Postgres.Host = viper.GetString("postgres.host")
	cfg.Postgres.Port = viper.GetString("postgres.port")
	cfg.Postgres.DB = viper.GetString("postgres.db")
	cfg.Postgres.User = viper.GetString("postgres.user")
	cfg.Postgres.Password = viper.GetString("postgres.password")

	cfg.Redis.Host = viper.GetString("redis.host")
	cfg.Redis.Port = viper.GetString("redis.port")
	cfg.Redis.Password = viper.GetString("redis.password")
	cfg.Redis.Expires = viper.GetDuration("redis.expires")

	cfg.WS.Port = viper.GetString("websocket.port")

	cfg.Mongo.Host = viper.GetString("mongo.host")
	cfg.Mongo.Port = viper.GetString("mongo.port")
	cfg.Mongo.User = viper.GetString("mongo.user")
	cfg.Mongo.Password = viper.GetString("mongo.password")
	cfg.Mongo.DB = viper.GetString("mongo.db")
}
