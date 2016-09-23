package message

import (
	"fmt"
	"net/http"

	"github.com/kkrs/di"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// DSTransport implements Transport by backing messages to Datastore. It has
// request lifetime because the field Context needs to be created for every
// request.
type DSTransport struct {
	Ctx context.Context
}

// Send persists the message to datastore.
func (tr DSTransport) Send(msg Message) error {
	key := datastore.NewIncompleteKey(tr.Ctx, "message",
		datastore.NewKey(tr.Ctx, "root", "root", 0, nil),
	)
	_, err := datastore.Put(tr.Ctx, key, &msg)
	return err
}

// List retrieves the first 10 messages from datastore.
func (tr DSTransport) List() ([]Message, error) {
	msgs := make([]Message, 0, 10)
	q := datastore.NewQuery("message").Ancestor(
		datastore.NewKey(tr.Ctx, "root", "root", 0, nil),
	)
	_, err := q.GetAll(tr.Ctx, &msgs)
	return msgs, err
}

// ListTransport implements Transport and stores messages in a slice. It is
// required to be a singleton so that the messages stored in it are not
// lost.
type ListTransport struct {
	msgs []Message
}

func (tr *ListTransport) Send(msg Message) error {
	tr.msgs = append(tr.msgs, msg)
	return nil
}

func (tr *ListTransport) List() ([]Message, error) {
	return tr.msgs, nil
}

// ReqFactory knows how to create Controllers and its dependencies.
type ReqFactory struct {
	af  AppFactory // access to singletons
	req *http.Request
}

func (fa ReqFactory) newTransport() Transport {
	switch fa.af.Env {
	case "e2e":
		return DSTransport{appengine.NewContext(fa.req)}
	case "int":
		return fa.af.ListTr
	default:
		panic(fmt.Sprintf("do not know how to make Transport for env %q", fa.af.Env))
	}
}

func (fa ReqFactory) NewController(label string) di.Controller {
	switch label {
	case "message":
		return MessageController{fa.newTransport()}
	default:
		panic(fmt.Sprintf("do not know how to make %q", label))
	}
}

// AppFactory contains singletons.
type AppFactory struct {
	Env    string
	ListTr *ListTransport
}

func (fa AppFactory) With(req *http.Request) di.RequestFactory {
	return ReqFactory{fa, req}
}
