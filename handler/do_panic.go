package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

type DoPanicHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var doPanicHandler = &model.HealthzResponse{}
	doPanicHandler.Message = "OK"
	panic("This is a panic example")
	// println(healthzHandler.Message)

	json.NewEncoder(w).Encode(doPanicHandler)
	return
}
