package middleware

import (
    "net/http"
    "time"
    
    log "github.com/sirupsen/logrus"
)

// Logger - это middleware для логирования HTTP-запросов.
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Записываем время начала для подсчета длительности
        start := time.Now()

        // Создаем логгер с базовыми полями из запроса
        logger := log.WithFields(log.Fields{
            "method":     r.Method,
            "path":       r.URL.Path,
            "request_id": r.Header.Get("X-Request-ID"),
        })

        // Логируем начало запроса
        logger.Info("Request started")

        // Оборачиваем ResponseWriter, чтобы перехватить код статуса и размер ответа
        lrw := newLoggingResponseWriter(w)

        // Передаем управление следующему обработчику в цепочке
        next.ServeHTTP(lrw, r)

        // Логируем завершение запроса с дополнительными метриками
        logger.WithFields(log.Fields{
            "status":     lrw.statusCode,
            "duration":   time.Since(start).String(),
            "duration_ms": time.Since(start).Milliseconds(),
        }).Info("Request completed")
    })
}

// --- Вспомогательные типы и методы для перехвата статуса и размера ответа ---
type loggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
    return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}