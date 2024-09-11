package api

import (
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	_ "recipe-management/internal/api/docs"
	"recipe-management/internal/api/handler"
	"recipe-management/internal/api/middleware"
	"recipe-management/internal/config"
	"recipe-management/internal/storage"

	"github.com/coocood/freecache"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Recipe Management System
// @version 1.0
// @description API for Recipe Management System
// @host localhost:8081
// @BasePath /recipes
// @schemes http
func NewRouter(cfg *config.Config, logger *slog.Logger, storage storage.IStorage) *gin.Engine {
	h := handler.NewHandler(storage, logger)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	targetURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatalf("error while parsing target URL: %v", err)
	}

	var cache = freecache.NewCache(cfg.CACHE_SIZE)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.Director = func(req *http.Request) {
		req.Header.Set("X-Forwarded-Host", req.RemoteAddr)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
	}

	router.Use(middleware.CacheAndLog(cache, logger, targetURL))

	router.POST("/recipes", h.Add)
	router.GET("/recipes/:id", h.Get)
	router.PUT("/recipes/:id", h.Update)
	router.DELETE("/recipes/:id", h.Delete)
	router.GET("/recipes", h.List)

	router.Any("/proxy/*any", gin.WrapH(proxy))
	log.Println(targetURL.String())
	log.Println(cfg.CACHE_SIZE, cfg.CACHE_TIME)

	return router
}
