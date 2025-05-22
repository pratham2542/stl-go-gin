package internals

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppContext interface {
	IsAppContext()
}

type HandlerFunc func(AppContext, *gin.Context) (any, error)

func HandlerWrapper(appCtx AppContext, handlerf HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := handlerf(appCtx, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		var statusCode int = http.StatusInternalServerError
		switch c.Request.Method {
		case http.MethodPut, http.MethodPatch:
			statusCode = http.StatusNoContent
		case http.MethodPost:
			statusCode = http.StatusCreated
		case http.MethodDelete:
			statusCode = http.StatusNoContent
		default:
			statusCode = http.StatusOK
		}

		switch v := res.(type) {
		case nil:
			c.Status(http.StatusNoContent)
		case gin.H:
			c.JSON(statusCode, v)
		case string:
			c.String(statusCode, v)
		default:
			c.JSON(statusCode, v)
		}
	}
}
