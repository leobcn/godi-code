package message_test

import (
	"net/http"
	"testing"

	. "github.com/kkrs/godi-code"
)

func TestSend(t *testing.T) {
	t.Logf("Scenario: Sending a message delivers it successfully")
	t.Log()

	server := "http://localhost:8080"
	msg := Message{"kkrs", "world", "hello"}
	// create request to send message
	req, desc := sendRequest(server, msg)
	resp, err := http.DefaultClient.Do(req)
	verify(t, desc, resp, err, http.StatusOK, nil)

	// create request to list all messages sent
	req, desc = listRequest(server)
	resp, err = http.DefaultClient.Do(req)
	// verify that it contains the one sent earlier
	verify(t, desc, resp, err, http.StatusOK, []Message{msg})
}
