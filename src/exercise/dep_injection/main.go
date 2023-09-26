package dep_injection

import (
	"fmt"
	"net/http"
)

func Main() {
	logger := LoggerAdapter(func(m string) { fmt.Println(m) })
	controller := NewController(
		logger,
		NewSimpleService(logger, NewSimpleDataStore()),
	)
	http.HandleFunc("/", controller.HandleGreetings)
	http.ListenAndServe(":8080", nil)
}
