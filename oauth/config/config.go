package config

import "os"

//getEnv
func GetEnv(name, defaultValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defaultValue
}
