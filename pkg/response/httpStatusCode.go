package response

import (
	"net/http"
	"strings"
)

func GetResponseMessage(code int) string {
	message := make(map[int]string)

	message[http.StatusOK] = "success"
	message[http.StatusCreated] = "created"
	message[http.StatusAccepted] = "accepted"
	message[http.StatusBadRequest] = "bad request"
	message[http.StatusUnauthorized] = "unauthorized"
	message[http.StatusForbidden] = "forbidden"
	message[http.StatusNotFound] = "not found"
	message[http.StatusTooManyRequests] = "too many request"
	message[http.StatusInternalServerError] = "internal server error"
	message[http.StatusServiceUnavailable] = "service unavailable"

	if _, ok := message[code]; !ok {
		return strings.ToLower(http.StatusText(code))
	}

	return message[code]
}
