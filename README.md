# Go Config Loader

A tiny Go library aimed to provide a simple method of of app configuration(s) from various sources and formats.

## Usage

`database.yaml`:

```yaml
database:
  host: localhost
  port: 5432
  user: root
  pass: root
```

`rabbitmq.json`:

```json
{
  "rabbitmq": {
    "host": "localhost",
    "port": 5432,
    "user": "root",
    "pass": "root",
    "vhost": "/foobar"
  }
}
```

```go
package main

import (
	"log"

	"github.com/ashep/go-cfgloader"
)

type Database struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	User string `json:"user" yaml:"user"`
	Pass string `json:"pass" yaml:"pass"`
}

type RabbitMQ struct {
	Host  string `json:"host" yaml:"host"`
	Port  int    `json:"port" yaml:"port"`
	User  string `json:"user" yaml:"user"`
	Pass  string `json:"pass" yaml:"pass"`
	VHost string `json:"vhost" yaml:"vhost"`
}

type Config struct {
	DB Database `json:"database" yaml:"database"`
	MQ RabbitMQ `json:"rabbitmq" yaml:"rabbitmq"`
}

func main() {
	cfg := Config{}

	if err := cfgloader.Load("database.yaml", &cfg, nil); err != nil {
		log.Panicf("failed to load config file: %s", err)
	}

	if err := cfgloader.Load("rabbitmq.json", &cfg, nil); err != nil {
		log.Panicf("failed to load config file: %s", err)
	}
}
```
