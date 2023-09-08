package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mileusna/useragent"
)

type osNameKey int

const (
	OSNameContextKey osNameKey = 12345
)

func GetDevice(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// HTTPリクエストからContextを取得
		ctx := r.Context()

		// Contextを更新し、OS名を格納
		ua := r.Header.Get("User-Agent")
		userAgent := useragent.Parse(ua)
		ctx = context.WithValue(ctx, OSNameContextKey, userAgent.OS)

		// 更新したContextを新しいhttp.Requestに設定
		r = r.WithContext(ctx)

		fmt.Println("ｒは")
		fmt.Println(r.Context().Value(OSNameContextKey).(string))

		// OS名を取得して表示
		osName, ok := ctx.Value(OSNameContextKey).(string)
		if !ok {
			// キーが存在しない場合のエラーハンドリング
			fmt.Println("OS name not found in context")
		} else {
			fmt.Println("OS:", osName)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
