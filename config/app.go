package config

import "gex-api/pkg/env"

type AppConfig struct {
	Host   string
	Name   string
	Mode   string
	Key    string
	Domain struct {
		SSL  bool
		Name string
	}
}

var App = &AppConfig{
	Host: env.Get("app.host", "localhost:3000").(string),
	Name: env.Get("app.name", "gex").(string),
	Mode: env.Get("app.mode", "debug").(string),
	Key:  env.Get("app.key", "kHH8AiNAdaf1e4QhXJYZ").(string),
	Domain: struct {
		SSL  bool
		Name string
	}{
		SSL:  env.Get("app.domain.ssl", false).(bool),
		Name: env.Get("app.domain.name", "domain.com").(string),
	},
}
