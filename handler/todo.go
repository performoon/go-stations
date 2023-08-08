package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//var healthzHandler = &model.HealthzResponse{}
	var createTODORequest = &model.CreateTODORequest{}
	var createTODOResponse = &model.CreateTODOResponse{}
	if r.Method == "POST" {
		json.NewDecoder(r.Body).Decode(createTODORequest)
		fmt.Println("Method=POST")
		if createTODORequest.Subject == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			// http.Error(w, "Bad Request: Invalid input", http.StatusBadRequest)
			// w.Header().Set("Content-Type", "application/json") // レスポンスのContent-Typeを設定
			// w.WriteHeader(http.StatusBadRequest)
			// json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: Invalid input"})
			return
			// w.Write()
		} else if createTODORequest.Description == "" {
			todo, err := h.svc.CreateTODO(r.Context(), createTODORequest.Subject, "")
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError) // サーバーエラーの場合のステータスコードを設定
				json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
				return
			}
			//createTODOResponse.TODO.Subject = h.
			fmt.Print("todo:")
			fmt.Print(todo)
			createTODOResponse.TODO = *todo
			json.NewEncoder(w).Encode(createTODOResponse)
		} else {
			todo, err := h.svc.CreateTODO(r.Context(), createTODORequest.Subject, createTODORequest.Description)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError) // サーバーエラーの場合のステータスコードを設定
				json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
				return
			}
			//createTODOResponse.TODO.Subject = h.
			createTODOResponse.TODO = *todo
			json.NewEncoder(w).Encode(createTODOResponse)
			return
		}
	}
	// println(healthzHandler.Message)

	//json.NewEncoder(w).Encode(healthzHandler)
	return
}
