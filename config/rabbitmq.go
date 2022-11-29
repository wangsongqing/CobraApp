package config

import (
	"CobraApp/pkg/config"
)

func init() {
	config.Add("rabbitmq", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("RABBITMQ_HOST", "127.0.0.1"),
			"port":     config.Env("RABBITMQ_PORT", "9200"),
			"user":     config.Env("RABBITMQ_USER", "guest"),
			"password": config.Env("RABBITMQ_PASSWORD", "guest"),
		}
	})
}
