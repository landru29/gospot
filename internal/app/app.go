// Package app is the main application.
package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/landru29/gospot/internal/server"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Controller is the application controller.
type Controller struct {
	conf    *Config
	log     logrus.FieldLogger
	apiList []server.Router
}

// Config is the configuration object.
type Config struct {
	ClientID     string `yaml:"clientId"      json:"clientId"`
	ClientSecret string `yaml:"clientSecret"  json:"clientSecret"`
	ClientBind   string `yaml:"clientBind"    json:"clientBind"`
	APIBind      string `yaml:"apiBind"       json:"apiBind"`
	Redirect     string `yaml:"redirect"      json:"redirect"`
	Login        string `yaml:"login"         json:"login"`
}

// LoadConfig loads a configuration from file.
func (c *Config) LoadConfig(confDir string, name string) error {
	for _, dir := range []string{".", confDir} {
		for ext, format := range map[string]string{
			"json": "json",
			"yaml": "yaml",
			"yml":  "yaml",
		} {
			file, err := os.Open(path.Clean(path.Join(dir, fmt.Sprintf("%s.%s", name, ext))))
			if err == nil {
				defer func() { _ = file.Close() }()

				switch format {
				case "json":
					return json.NewDecoder(file).Decode(c)

				case "yaml":
					return yaml.NewDecoder(file).Decode(c)
				}
			}
		}
	}

	return fmt.Errorf("could not find any suitable configuration")
}

// New creates a controller.
func New(conf *Config, log logrus.FieldLogger, apiList ...server.Router) (*Controller, error) {
	return &Controller{
		conf:    conf,
		log:     log,
		apiList: apiList,
	}, nil
}

// Start launch the application.
func (c *Controller) Start(ctx context.Context, quit chan os.Signal) {
	cancelFuncs := make([]context.CancelFunc, len(c.apiList))

	for idx, apiServer := range c.apiList {
		cancellableCtx, cancel := context.WithCancel(ctx)
		cancelFuncs[idx] = cancel

		server := apiServer

		c.log.WithField("idx", idx).Info("starting service")

		go func() {
			server.Start(cancellableCtx)
		}()
	}

	<-quit

	for idx, cancellable := range cancelFuncs {
		c.log.WithField("idx", idx).Info("stopping service")
		cancellable()
	}
}
