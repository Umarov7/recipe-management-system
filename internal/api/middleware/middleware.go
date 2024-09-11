package middleware

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/coocood/freecache"
	"github.com/gin-gonic/gin"
)

func CacheAndLog(cache *freecache.Cache, logger *slog.Logger, targetURL *url.URL) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.Path

		data, err := cache.Get([]byte(key))
		if err == nil {
			logger.Info("Cache hit", slog.String("key", key))
			c.Data(http.StatusOK, "application/json", data)
			c.Abort()
			return
		}
		logger.Info("Cache miss", slog.String("key", key))

		start := time.Now()
		writer := &responseWriter{c.Writer, bytes.NewBuffer([]byte{})}
		c.Writer = writer

		c.Next()

		if c.Writer.Status() == http.StatusOK {
			data := writer.Buffer.Bytes()
			if err := cache.Set([]byte(key), data, 90); err != nil {
				logger.Error("failed to cache response", slog.String("key", key), slog.String("error", err.Error()))
			} else {
				logger.Info("Cache set", slog.String("key", key))
			}
		}

		duration := time.Since(start)
		clientIP := c.GetHeader("X-Forwarded-For")
		if clientIP == "" {
			clientIP = c.ClientIP()
		}

		backendServer := targetURL.String()

		logger.Info("Request Details",
			slog.String("client_ip", clientIP),
			slog.String("backend_server", backendServer),
			slog.Duration("duration", duration),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
		)
	}
}

type responseWriter struct {
	gin.ResponseWriter
	Buffer *bytes.Buffer
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.Buffer.Write(data)
	return rw.ResponseWriter.Write(data)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
}
