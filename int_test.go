// +build int

package message_test

import (
	"net/http/httptest"
	"testing"

	. "github.com/kkrs/godi-code"
)

func TestSend(t *testing.T) {
	transport := Setup(
		AppFactory{"int", &ListTransport{}}, []Registration{
			{MessageController{}, "message"},
		})

	server := httptest.NewServer(transport)
	testSend(t, server.URL)
}
