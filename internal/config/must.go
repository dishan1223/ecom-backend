package config

import "os"

func MustLoadEnv(a string) string {
	d := os.Getenv(a)
	if d == "" {
		panic("environment variable " + a + " not set")
	}

	return d
}
