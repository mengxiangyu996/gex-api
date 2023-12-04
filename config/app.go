package config

import "breeze-api/pkg/env"

type AppConfig struct {
	Name string
	Host string
	Port int
	Key  string
}

var App = &AppConfig{
	Name: env.Get("app.name", "breeze-api").(string),
	Host: env.Get("app.host", "127.0.0.1").(string),
	Port: int(env.Get("app.port", 3000).(float64)),
	Key:  env.Get("app.key", "breeze-api").(string),
}
