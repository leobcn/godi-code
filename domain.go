package message

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kkrs/di"
	"github.com/kkrs/di/router"
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

// MessageController handles requests to send and list messages.
type MessageController struct {
	Transport Transport // dependency injected
}

// MessageController would like POST:/api/messages to be dispatched to Send.
func (MessageController) Bindings() []di.Binding {
	return []di.Binding{
		{"POST", APIPath, "Send"},
		{"GET", SpyPath, "List"},
	}
}

// Send processes the request and delegates the task of sending the message to
// Transport.
func (ct MessageController) Send(rw http.ResponseWriter, req *http.Request) {
	var msg Message
	if err := Unmarshal(req.Body, &msg); err != nil {
		HTTPError(
			rw,
			http.StatusInternalServerError,
			fmt.Errorf("error reading request: %s", err),
		)
	}

	if err := ct.Transport.Send(msg); err != nil {
		HTTPError(
			rw,
			http.StatusInternalServerError,
			fmt.Errorf("error sending message: %s", err),
		)
	}
	rw.WriteHeader(http.StatusOK)
}

// List processes the request and delegates the task of listing messages to
// Transport.
func (ct MessageController) List(rw http.ResponseWriter, req *http.Request) {
	msgs, err := ct.Transport.List()
	if err != nil {
		HTTPError(
			rw,
			http.StatusInternalServerError,
			fmt.Errorf("error getting messages: %s", err),
		)
		return
	}

	data, err := json.Marshal(msgs)
	if err != nil {
		HTTPError(
			rw,
			http.StatusInternalServerError,
			fmt.Errorf("error marshalling results: %s", err),
		)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

// Registration is used to pass arguments to Setup
type Registration struct {
	Ctrl  di.Controller
	Label string
}

func Setup(af di.ApplicationFactory, regs []Registration) di.Router {
	router := router.New()
	dispatcher := di.New("messageService", router, af)
	for _, r := range regs {
		if err := dispatcher.Register(r.Ctrl, r.Label); err != nil {
			panic(err)
		}
	}
	return router
}
