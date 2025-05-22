package internals

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type AppContext interface {
	IsAppContext()
}

type HandlerFunc func(AppContext, *gin.Context) (any, error)

func HandlerWrapper(appCtx AppContext, handlerf HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp Response
		defer func() {
			switch exception := recover(); exception {
			case nil:
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

				if resp.Err != nil {
					c.JSON(http.StatusInternalServerError, ErrorResponse{Error: resp.Err.Error()})
					return
				}

				switch v := resp.Res.(type) {
				case nil:
					c.Status(http.StatusNoContent)
				case gin.H:
					c.JSON(statusCode, v)
				case string:
					c.String(statusCode, v)
				default:
					c.JSON(statusCode, v)
				}
			default:
				stack := debug.Stack()
				fmt.Printf("Panic recovered: %v\nStack trace:\n%s\n", exception, stack)
				c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal Server Error"})
			}
		}()

		res, err := handlerf(appCtx, c)
		resp = Response{
			Res: res,
			Err: err,
		}
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Res any
	Err error
}
