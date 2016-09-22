package message

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func HTTPError(rw http.ResponseWriter, status int, err error) {
	http.Error(rw, fmt.Sprintf(`{"error": "%s"}`, err.Error()), status)
}

func Unmarshal(body io.Reader, dst interface{}) error {
	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(payload, dst); err != nil {
		return err
	}
	return nil
}

var (
	APIPath = "/api/messages"
	SpyPath = "/spy/messages"
)

type Message struct {
	From    string
	To      string
	Message string
}

// Transport represents the ability to send a Message.
type Transport interface {
	Send(Message) error
	List() ([]Message, error) // List messages sent
}
