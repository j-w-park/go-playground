package dep_injection

import (
	"net/http"
)

type Controller struct {
	logger  Logger
	service Service
}

func (c Controller) HandleGreetings(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if message, err := c.service.SayHelloTo(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		c.logger.Log("HandleGreetings: " + message)
		w.Write([]byte(message))
	}
}

func NewController(logger Logger, service Service) Controller {
	return Controller{logger: logger, service: service}
}
