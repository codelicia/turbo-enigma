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

func (h *HealthCheck) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "It is alive!")
}
