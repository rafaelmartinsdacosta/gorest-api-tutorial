// +build e2e

package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get(os.Getenv("SERVER_TEST") + "/api/health")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
