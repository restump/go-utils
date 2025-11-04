package utils

import "os"

func LookupEnvWithDefault(k string, dv string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return dv
}
