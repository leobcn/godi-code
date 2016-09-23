package message_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	. "github.com/kkrs/godi-code"
)

func sendRequest(address string, msg Message) (*http.Request, string) {
	urlStr := APIPath
	if len(address) > 0 {
		urlStr = address + APIPath
	}
	body, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	return req, fmt.Sprintf("Request POST, %s with body '%s'", APIPath, string(body))
}

func listRequest(address string) (*http.Request, string) {
	urlStr := SpyPath
	if len(address) > 0 {
		urlStr = address + SpyPath
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}
	return req, fmt.Sprintf("Request GET, %s", SpyPath)
}

// verify resp against expected status, body
func verify(t *testing.T, desc string, resp *http.Response, err error, status int, body interface{}) {
	t.Log(desc, " should succeed")
	if err != nil {
		t.Fatalf("got error '%s'", err)
	}
	t.Logf("and response should have")
	t.Logf("\tstatus '%s'", http.StatusText(status))
	if resp.StatusCode != status {
		t.Fatalf("got status '%s', but expected '%s'", http.StatusText(resp.StatusCode), http.StatusText(status))
	}
	if body != nil {
		t.Logf("\tbody that that unmarshals to %#v", body)
		got := reflect.New(reflect.TypeOf(body)).Interface()
		if err := Unmarshal(resp.Body, got); err != nil {
			t.Fatalf("got error '%s'", err)
		}
		if !reflect.DeepEqual(reflect.ValueOf(got).Elem().Interface(), body) {
			t.Fatalf("got %+v", got)
		}
	}
}

func testSend(t *testing.T, server string) {
	t.Logf("Scenario: Sending a message delivers it successfully")
	t.Log()
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
