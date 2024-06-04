package ports

import "net/http"

type HTTPContext interface {
	JSON(int, interface{})
	Request() *http.Request
}
