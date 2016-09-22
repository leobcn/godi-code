// +build appengine

package message

import "net/http"

func init() {
	http.Handle(APIPath, http.HandlerFunc(Send))
	http.Handle(SpyPath, http.HandlerFunc(List))
}
