package config

import "os"

func getEnv(name, defaultValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defaultValue
}
