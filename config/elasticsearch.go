package config

import (
	"CobraApp/pkg/config"
)

func init() {

	config.Add("elasticsearch", func() map[string]interface{} {
		return map[string]interface{}{
			"host": config.Env("ES_HOST", "127.0.0.1"),
			"port": config.Env("ES_PORT", "9200"),
		}
	})
}
