package health

import (
	"test-go/internals"

	"github.com/gin-gonic/gin"
)

func HealthController(appCtx internals.AppContext, c *gin.Context) (any, error) {
	user := map[string]string{"id": "1", "name": "Alice"}
	return gin.H{"user": user}, nil
}

func DeleteController(appCtx internals.AppContext, c *gin.Context) (any, error) {
	return gin.H{"deleted": true}, nil
}

func CreateController(appCtx internals.AppContext, c *gin.Context) (any, error) {
	return gin.H{"created": true}, nil
}

func UpdateController(appCtx internals.AppContext, c *gin.Context) (any, error) {
	return gin.H{"Updated": true}, nil
}
