package util

import (
	"os"
)

// Gets the value for the given environment variable, or sets the given default
func GetEnvVarOrDefault(env string, def string) string {
	envVar := os.Getenv(env)
	if len(envVar) == 0 {
		return def
	}
	return envVar
}
