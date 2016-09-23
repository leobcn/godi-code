package message

import (
	"net/http"

	. "github.com/kkrs/godi-code"
)

func init() {
	http.Handle(APIPath, http.HandlerFunc(Send))
	http.Handle(SpyPath, http.HandlerFunc(List))
}
