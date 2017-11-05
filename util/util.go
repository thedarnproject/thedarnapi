package util

import (
	"os"
)

// Gets the value for the given environment variable, or sets the given default
func GetEnvVarOrDefault(env, def string) string {
	envVar, found := os.LookupEnv(env)
	if found {
		return envVar
	}
	return def
}
