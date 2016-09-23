// +build e2e

package message_test

import (
	"testing"
)

func TestSend(t *testing.T) {
	testSend(t, "http://localhost:8080")
}
