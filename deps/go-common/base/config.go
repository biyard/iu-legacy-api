package base

import (
	"flag"
	"io/ioutil"
	"net/http"

	"biyard.co/common/logger"
	"gopkg.in/yaml.v2"
)

type ConfigGetter interface {
	Get() *Config
}

type (
	Env        string
	ServerType string
)

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvProd  Env = "prod"

	ServerTypeGin        ServerType = "gin"
	ServerTypeLambda     ServerType = "lambda"
	ServerTypeLambdaTask ServerType = "lambda-task"
	ServerTypeWebWorker  ServerType = "web-worker"
)

type Config struct {
	Environment Env `yaml:"env"`

	Server          ServerConfig         `yaml:"server"`
	Watch           WatchConfig          `yaml:"watch"`
	Logging         logger.LoggingConfig `yaml:"logging"`
	CORS            CORSConfig           `yaml:"cors"`
	Database        DatabaseConfig       `yaml:"database"`
	DiscordDatabase DatabaseConfig       `yaml:"discordDatabase"`
	AWS             AWSConfig            `yaml:"aws"`
	Redis           RedisConfig          `yaml:"redis"`
}

type AWSConfig struct {
	SecretKey   string `yaml:"secretKey"`
	AccessKeyID string `yaml:"accessKeyId"`
	Region      string `yaml:"region"`
}

type CORSConfig struct {
	Enabled         bool     `yaml:"enabled"`
	Origins         []string `yaml:"origins"`
	AllowAllOrigins bool     `yaml:"allowAllOrigins"`
	Headers         []string `yaml:"headers"`
	Methods         []string `yaml:"methods"`
	CredentialMode  bool     `yaml:"credentialMode"`
}

type DatabaseConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Type     string `yaml:"type"`
	Endpoint string `yaml:"endpoint"`
	Name     string `yaml:"name"`
}

type RedisConfig struct {
	Enabled         bool   `yaml:"enabled"`
	Single          bool   `yaml:"single"`
	TLS             bool   `yaml:"tls"`
	ServiceID       string `yaml:"serviceId"`
	Endpoint        string `yaml:"endpoint"`
	CacheExpiration string `yaml:"cacheExpiration"`
	MinIdleConn     int    `yaml:"minIdleConn"`
	TTL             int    `yaml:"ttl"`
	PoolSize        int    `yaml:"poolSize"`
	PoolTimeout     string `yaml:"poolTimeout"`
	ReadTimeout     string `yaml:"readTimeout"`
	WriteTimeout    string `yaml:"writeTimeout"`
	MaxRetries      int    `yaml:"maxRetries"`
	MinRetryBackoff string `yaml:"minRetryBackoff"`
	MaxRetryBackoff string `yaml:"maxRetryBackoff"`
	ReadOnly        bool   `yaml:"readOnly"`
	RandomRoute     bool   `yaml:"randomRoute"`
}

func (r *Config) Get() *Config {
	return r
}

type WatchConfig struct {
	Enabled bool `yaml:"enabled"`
}

type ServerConfig struct {
	Type ServerType `yaml:"type"`
	Port int        `yaml:"port"`
}

func DefaultConfig() Config {
	return Config{
		Environment: EnvLocal,
		Server: ServerConfig{
			Type: ServerTypeGin,
			Port: 3000,
		},
		Watch: WatchConfig{
			Enabled: true,
		},
		Logging: logger.LoggingConfig{
			Level:    "info",
			Encoding: "json",
			LogRotate: logger.LogRotateConfig{
				Enabled:      false,
				LogPath:      "logs/app.log",
				ErrorLogPath: "logs/error.log",
				MaxSize:      100,
				MaxBackups:   10,
				MaxAge:       30,
			},
		},
		CORS: CORSConfig{
			Enabled:        false,
			Origins:        []string{"http://localhost:5000"},
			Headers:        []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
			Methods:        []string{http.MethodPost, http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch},
			CredentialMode: true,
		},
		Database: DatabaseConfig{
			Enabled:  true,
			Type:     "dynamodb",
			Endpoint: "",
			Name:     "local",
		},
		AWS: AWSConfig{
			SecretKey:   "",
			AccessKeyID: "",
			Region:      "ap-northeast-2",
		},
	}
}

// Load loads config from a file. If the `-config` flag is not provided, it will
// load `config.yaml` by default.
func Load(cfg any) {
	conf := flag.String("config", "config.yaml", "indicates a config file")
	flag.Parse()

	cdata, err := ioutil.ReadFile(*conf)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(cdata, cfg); err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(cdata, cfg.(ConfigGetter).Get()); err != nil {
		panic(err)
	}
}
