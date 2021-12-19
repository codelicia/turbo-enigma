package handler

import (
	"fmt"
	"net/http"
)

type HealthCheck struct {
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

//revive:disable:unused-parameter Needed from the interface
func (h *HealthCheck) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "It is alive!")
}
