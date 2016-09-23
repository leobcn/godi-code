package message

import (
	"net/http"

	. "github.com/kkrs/godi-code"
)

func init() {
	router := Setup(AppFactory{"e2e", nil}, []Registration{
		{MessageController{}, "message"},
	})
	http.Handle("/", router)
}
