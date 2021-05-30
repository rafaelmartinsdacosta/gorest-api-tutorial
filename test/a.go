// +build e2e

package test

import "os"

func SetupEnv() {
	host := os.Getenv("SERVER_TEST")
	if len(host) != 21 {
		os.Setenv("SERVER_TEST", "http://localhost:8080")
	}
}

func init() {
	SetupEnv()
}
