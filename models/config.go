package models

import (
	"crypto/sha256"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	configPath = "config/config.toml"
)

type duration time.Duration

// Config struct
type Config struct {
	SQLDataBase SQLDataBase `toml:"SQLDataBase"`
	ServerOpt   ServerOpt   `toml:"ServerOpt"`
	HashSum     []byte
}

func (d *duration) UnmarshalText(text []byte) error {
	temp, err := time.ParseDuration(string(text))
	*d = duration(temp)
	return err
}

// ServerOpt struct
type ServerOpt struct {
	ReadTimeout  duration
	WriteTimeout duration
	IdleTimeout  duration
}

// LoadConfig from path
func LoadConfig(c *Config) {
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		return
	}
	c.SQLDataBase.UserID = getCredential("/etc/scrt/service-template/sqlUser")
	c.SQLDataBase.Password = getCredential("/etc/scrt/service-template/sqlPassword")

	c.HashSum = GetHashSum()
}

func getCredential(path string) string {
	c, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(c))
}

// GetHashSum of config file
func GetHashSum() []byte {
	paths := []string{
		configPath,
		"/etc/scrt/service-template/sqlUser",
		"/etc/scrt/service-template/sqlPassword",
	}
	h := sha256.New()

	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return nil
		}
		defer f.Close()
		if _, err = io.Copy(h, f); err != nil {
			return nil
		}
	}

	return h.Sum(nil)
}
