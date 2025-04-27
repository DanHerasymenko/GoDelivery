package response

import "github.com/gin-gonic/gin"

// AbortWithErrorJson is a utility function to abort the request with a JSON error response
// and add error to context to display in logging middleware.
func AbortWithErrorJson(ctx *gin.Context, statusCode int, err error) {

	ctx.Error(err)

	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"error": err.Error(),
	})

}
