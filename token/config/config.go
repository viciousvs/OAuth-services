package config

import (
	"os"
	"strconv"
)

//getEnv
func getEnv(name, defaultValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	if value, exists := os.LookupEnv(name); exists {
		val, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return val
	}
	return defaultValue
}
