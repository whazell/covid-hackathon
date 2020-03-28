package covid

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Config struct {
	Bind     string         `yaml:"bind"`
	Database DatabaseConfig `yaml:"database"`
	Log      LogConfig      `yaml:"log"`
}

type LogConfig struct {

	// Type can be either: stderr, file
	Type string `yaml:"type"`

	// Level that should be logged (see logrus)
	Level string `yaml:"level"`

	// Set logs to use JSON format or not
	Json bool `yaml:"json"`

	// Settings for when using file type
	EnableRolling bool   `yaml:"enableRolling"`
	MaxSize       int    `yaml:"maxSize"`
	MaxDays       int    `yaml:"maxDays"`
	FileName      string `yaml:"fileName"`

	// Devmode turns on logrus' pretty print mode
	DevMode bool `yaml:"dev"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func LoadConfigFile(file string) (Config, error) {
	if _, err := os.Stat(file); err != nil {
		return Config{}, err
	}
	cfgData, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	return LoadConfig(cfgData)
}

func LoadConfig(data []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}
	return c, c.Validate()
}

func (c Config) Validate() error {
	if c.Bind == "" {
		return errors.New("Must provide port to expose application on")
	}

	if c.Bind == ":80" || c.Bind == ":443" {
		log.Warn("Binding to :80 or :443 requires root priviledges")
	}

	if err := c.Database.Validate(); err != nil {
		return err
	}
	return c.Log.Validate()
}

func (c DatabaseConfig) Validate() error {
	if c.Name == "" {
		return errors.New("Must provide database name")
	}
	if c.Port == "" {
		log.Warn("Database port not specified, defaulting to 3306")
		c.Port = "3306"
	}
	if c.Host == "" {
		log.Warn("Database host not provided, defaulting to 127.0.0.1")
		c.Host = "127.0.0.1"
	}
	if c.User == "" {
		return errors.New("Must provide database user")
	}
	if c.Password == "" {
		return errors.New("Must provide database password")
	}
	return nil
}

func (c LogConfig) Validate() error {
	if c.Type != "file" && c.Type != "stderr" {
		return errors.New("Log type must be one of: ['file', 'stderr']")
	}
	if c.Type == "file" && c.FileName == "" {
		return errors.New("Must provide file name")
	}
	return nil
}

// SetupLogger will configure the underlying log package (currently logrus) to match the
// options specified. This should be called after the LoadConfig() functions.
func (l LogConfig) SetupLogger() error {
	if l.Type == "file" {
		if l.EnableRolling {
			log.SetOutput(&lumberjack.Logger{
				Filename: l.FileName,
				MaxSize:  l.MaxSize,
				MaxAge:   l.MaxDays,
				Compress: false,
			})
		} else {
			flags := os.O_APPEND | os.O_WRONLY | os.O_CREATE
			f, err := os.OpenFile(l.FileName, flags, 0600)
			if err != nil {
				return err
			}
			log.SetOutput(f)
		}
	}

	if l.Json {
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		})
	} else {
		if !l.DevMode {
			log.SetFormatter(&log.TextFormatter{
				TimestampFormat: time.RFC3339Nano,
				DisableColors:   true,
			})
		}
	}

	switch strings.ToUpper(l.Level) {
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	return nil
}
