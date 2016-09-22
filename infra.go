package message

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// Send wraps around DSTransport.Send and responds to POST:/api/messages
func Send(rw http.ResponseWriter, req *http.Request) {
	var msg Message
	err := Unmarshal(req.Body, &msg)
	if err != nil {
		HTTPError(
			rw,
			http.StatusInternalServerError,
			fmt.Errorf("error reading request: %s", err),
		)
	}

	err = DSTransport{appengine.NewContext(req)}.Send(msg)
	if err != nil {
		HTTPError(
			rw,
			http.StatusInternalServerError,
			fmt.Errorf("error sending message: %s", err),
		)
	}
	rw.WriteHeader(http.StatusOK)
}

// List wraps around DSTransport.List and responds to GET:/spy/messages
func List(rw http.ResponseWriter, req *http.Request) {
	msgs, err := DSTransport{appengine.NewContext(req)}.List()
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
