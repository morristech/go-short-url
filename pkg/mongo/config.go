package mongo

import "fmt"

type Config struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Options  string
}

func (config *Config) uri() string {

	uri := "mongodb://"

	// append username/password to uri
	if config.Username != "" {
		uri += fmt.Sprintf("%s:%s@", config.Username, config.Password)
	}

	// append port
	if config.Port != "" {
		uri += fmt.Sprintf(":%s", config.Port)
	}

	// append database
	uri += fmt.Sprintf("/%s", config.Database)

	return uri
}
