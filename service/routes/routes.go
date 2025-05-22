package routes

import (
	"fmt"
	"net/http"
	"os"
	"test-go/config"
	"test-go/internals"
	"test-go/service/api/controllers/health"

	"github.com/gin-gonic/gin"
)

type Server struct {
	appContext *config.AppContext
	router     *gin.Engine
}

func NewServer(appContext *config.AppContext) *Server {
	return &Server{
		appContext: appContext,
		router:     gin.New(),
	}
}

func (s *Server) Start() *http.Server {
	server := &http.Server{
		Addr:    ":" + s.appContext.GetPort(),
		Handler: s.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	return server
}

// App routes

func (s *Server) AddRoutes() *Server {
	baseUrl := s.router.Group("/apis")
	v1 := baseUrl.Group("/v1")
	v1.GET("/test", internals.HandlerWrapper(s.appContext, health.HealthController))
	v1.DELETE("/delete", internals.HandlerWrapper(s.appContext, health.DeleteController))
	v1.POST("/create", internals.HandlerWrapper(s.appContext, health.DeleteController))
	v1.PATCH("/updated/patch", internals.HandlerWrapper(s.appContext, health.UpdateController))
	v1.PUT("/updated/put", internals.HandlerWrapper(s.appContext, health.UpdateController))

	return s
}
