package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	healthHandler := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", healthHandler.ServeHTTP)

	todoHandler := handler.NewTODOHandler(service.NewTODOService(todoDB))
	mux.HandleFunc("/todos", todoHandler.ServeHTTP)

	doPanicHandler := handler.NewDoPanicHandler()
	recoveryMiddleware := middleware.Recovery
	mux.Handle("/do-panic", recoveryMiddleware(doPanicHandler))

	getDeviceMiddleware := middleware.GetDevice
	mux.Handle("/get-device", getDeviceMiddleware(todoHandler))

	requestInfoOutput := middleware.RequestInfoOutput
	mux.Handle("/request-info", getDeviceMiddleware(requestInfoOutput(todoHandler)))

	return mux
}
