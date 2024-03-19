package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s is required of (type %s)", name, typ)
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{"message": msg, "status": code})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s success full", op),
		"data":    data,
	})
}
