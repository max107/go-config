# install

```sh
go get github.com/max107/go-config
```

# usage

```go
package main

type kafkaCfg struct {
	Brokers       []string `env:"KAFKA_BROKERS"`
	User          string   `env:"KAFKA_USER"`
	Password      string   `env:"KAFKA_PASSWORD"`
	ConsumerGroup string   `env:"KAFKA_CONSUMER_GROUP"`
	Topics        []string `env:"KAFKA_TOPICS"`
	ClientID      string   `env:"KAFKA_CLIENT_ID"`
	GroupID       string   `env:"KAFKA_GROUP_ID"`
}

type redisCfg struct {
	RedisHost     string `env:"REDIS_HOST"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDatabase int    `env:"REDIS_DATABASE"`
}

var kafkaConfig = new(kafkaCfg)
var redisConfig = new(redisCfg)

func init() {
	config.Load([]string{".env", ".env.local"}, c, rc)
}

func main() {
    // ...
}
```
