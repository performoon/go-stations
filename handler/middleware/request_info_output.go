package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type accessTime int

type AccessLogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

const (
	accessTimeContextKey accessTime = iota
)

func RequestInfoOutput(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		beforAccessTime := time.Now()

		// リクエストにアクセス日時を追加
		ctx := r.Context()
		ctx = context.WithValue(ctx, accessTimeContextKey, beforAccessTime)
		r = r.WithContext(ctx)

		// ハンドラーを呼び出し、処理を続行
		h.ServeHTTP(w, r)

		afterAccessTime := time.Now()
		latency := afterAccessTime.Sub(beforAccessTime).Milliseconds()

		path := r.URL.Path
		fmt.Println("ｒは")
		fmt.Println(r.Context().Value(OSNameContextKey).(string))
		os := r.Context().Value(OSNameContextKey).(string)

		accessLog := AccessLogEntry{
			Timestamp: beforAccessTime,
			Latency:   latency,
			Path:      path,
			OS:        os,
		}

		jsonData, err := json.Marshal(accessLog)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println(string(jsonData))
	}
	return http.HandlerFunc(fn)
}
