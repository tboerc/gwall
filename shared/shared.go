package shared

import "os"

func Getenv(key string, fallback string) string {
	env, exist := os.LookupEnv(key)
	if exist {
		return env
	}
	return fallback
}
