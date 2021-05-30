// +build e2e

package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// TestGetComments - Unit test for get comments
func TestGetComments(t *testing.T) {
	fmt.Println("Running E2E test for get comments endpoint")

	client := resty.New()
	resp, err := client.R().Get(os.Getenv("SERVER_TEST") + "/api/comment")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

// TestPostComments - Unit test for post comment
func TestPostComments(t *testing.T) {
	fmt.Println("Running E2E test for post comment endpoint")

	client := resty.New()

	resp, err := client.R().
		SetBody(
			`{
			"slug": "/go-test",
			"body": "Unit test",
			"author": "test"
		}`).
		Post(os.Getenv("SERVER_TEST") + "/api/comment")

	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())
}
