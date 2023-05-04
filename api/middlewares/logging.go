package middlewares

import (
	"log"
	"net/http"

	"github.com/watarui/go-intermediate-practice/common"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{
		ResponseWriter: w,
		code:           http.StatusOK,
	}
}

func (w *resLoggingWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
