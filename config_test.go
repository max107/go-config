package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type firstCfg struct {
	NestedString string   `env:"NESTED_STRING"`
	AuthURL      string   `env:"AUTH_URL"`
	Strings      []string `env:"STRINGS"`
}

type secondCfg struct {
	Test1 bool `env:"BOOL_TEST1"`
	Test2 bool `env:"BOOL_TEST2"`
	Test3 bool `env:"BOOL_TEST3"`
	Test4 bool `env:"BOOL_TEST4"`
}

func TestLoadConfig(t *testing.T) {
	cfg := &firstCfg{}
	loadStruct(cfg, ".env.test")

	require.Equal(t, "qwe:8080", cfg.AuthURL)
	require.Equal(t, "foobar", cfg.NestedString)
}

func TestLoadConfigNoFiles(t *testing.T) {
	cfg := &firstCfg{}
	loadStruct(cfg, ".env.test")

	require.Equal(t, "qwe:8080", cfg.AuthURL)
	require.Equal(t, "foobar", cfg.NestedString)
	require.Equal(t, []string{"foo", "bar"}, cfg.Strings)
}

func TestLoadConfigs(t *testing.T) {
	cfg1 := &firstCfg{}
	cfg2 := &secondCfg{}

	Load([]string{".env", ".env.test"}, cfg1, cfg2)

	require.Equal(t, "qwe:8080", cfg1.AuthURL)
	require.Equal(t, "foobar", cfg1.NestedString)

	require.True(t, cfg2.Test1)
	require.False(t, cfg2.Test2)
	require.False(t, cfg2.Test3)
	require.False(t, cfg2.Test4)
}
