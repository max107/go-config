package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func loadStruct(cfg interface{}, fileNames ...string) {
	var valid []string
	for _, f := range fileNames {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			continue
		}
		valid = append(valid, f)
	}
	_ = godotenv.Overload(valid...)
	_ = env.Parse(cfg)
	return
}

// Load load env files into multiple structs
func Load(files []string, cfg ...interface{}) {
	for _, i := range cfg {
		loadStruct(i, files...)
	}
}
