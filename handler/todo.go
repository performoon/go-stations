package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

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

	//fmt.Println("start")
	switch r.Method {
	case "POST":
		var createTODORequest = &model.CreateTODORequest{}
		var createTODOResponse = &model.CreateTODOResponse{}
		json.NewDecoder(r.Body).Decode(createTODORequest)
		//fmt.Println("Method=POST")
		if createTODORequest.Subject == "" {
			// w.Header().Set("Content-Type", "text/plain")
			// w.WriteHeader(http.StatusBadRequest)
			// http.Error(w, "Bad Request: Invalid input", http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json") // レスポンスのContent-Typeを設定
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: Invalid input"})
			return
			// w.Write()
		} else if createTODORequest.Description == "" {
			fmt.Println("0")
			todo, err := h.svc.CreateTODO(r.Context(), createTODORequest.Subject, "")
			fmt.Println("1")
			if err != nil {
				fmt.Println("2")
				w.Header().Set("Content-Type", "application/json")
				fmt.Println("a")
				w.WriteHeader(http.StatusInternalServerError) // サーバーエラーの場合のステータスコードを設定
				fmt.Println("b")
				json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
				fmt.Println("c")
				return
			}
			fmt.Print("todo:")
			fmt.Print(todo.ID)
			fmt.Print("insertID type : ")
			fmt.Println(reflect.TypeOf(todo.ID))
			createTODOResponse.TODO = *todo
			fmt.Println("d")
			json.NewEncoder(w).Encode(createTODOResponse)
			fmt.Println("e")
			return
		} else {
			todo, err := h.svc.CreateTODO(r.Context(), createTODORequest.Subject, createTODORequest.Description)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError) // サーバーエラーの場合のステータスコードを設定
				json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
				//return
			}
			//createTODOResponse.TODO.Subject = h.
			createTODOResponse.TODO = *todo
			json.NewEncoder(w).Encode(createTODOResponse)
			return
		}
	case "PUT":
		var updateTODORequest = &model.UpdateTODORequest{}
		var updateTODOResponse = &model.UpdateTODOResponse{}
		json.NewDecoder(r.Body).Decode(updateTODORequest)
		//fmt.Println("Method=POST")
		if updateTODORequest.ID == 0 || updateTODORequest.Subject == "" {
			w.Header().Set("Content-Type", "application/json") // レスポンスのContent-Typeを設定
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request: Invalid input"})
		} else {
			todo, err := h.svc.UpdateTODO(r.Context(), updateTODORequest.ID, updateTODORequest.Subject, updateTODORequest.Description)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError) // サーバーエラーの場合のステータスコードを設定
				json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
				//return
			}
			//createTODOResponse.TODO.Subject = h.
			updateTODOResponse.TODO = *todo
			json.NewEncoder(w).Encode(updateTODOResponse)
			return
		}
	case "GET":
		var readTODORequest = &model.ReadTODORequest{}
		var readTODOResponse = &model.ReadTODOResponse{}

		//fmt.Println(r.URL.Query())
		getQuery := r.URL.Query()
		fmt.Println(getQuery)
		if getQuery.Has("prev_id") {
			getPrevID, err := strconv.Atoi(getQuery.Get("prev_id"))
			if err != nil {
				fmt.Print("prev_id err : ")
				fmt.Println(err)
			}
			readTODORequest.PrevID = int64(getPrevID)
		} else {
			readTODORequest.PrevID = 0
			fmt.Println("ID無いよ")
		}
		if getQuery.Has(("size")) {
			getSize, err := strconv.Atoi(getQuery.Get("size"))
			if err != nil {
				fmt.Print("size err : ")
				fmt.Println(err)
			}
			readTODORequest.Size = int64(getSize)
		} else {
			readTODORequest.Size = 5
			fmt.Println("size無いよ")
		}

		todos, err := h.svc.ReadTODO(r.Context(), readTODORequest.PrevID, readTODORequest.Size)
		if err != nil {
			fmt.Print("ReadTODO err : ")
			fmt.Println(err)
		}

		readTODOResponse.TODOs = todos

		json.NewEncoder(w).Encode(readTODOResponse)

		return
	}

	fmt.Println("どこも通ってない！？")
	return

}
