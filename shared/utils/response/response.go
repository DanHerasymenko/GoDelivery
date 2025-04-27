package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// AbortWithErrorJson is a utility function to abort the request
// and add error to context to display in logging middleware.
func AbortWithError(ctx *gin.Context, statusCode int, err error) {

	err = fmt.Errorf("%d: %w", statusCode, err)
	ctx.Error(err)
	ctx.AbortWithStatus(statusCode)

}
