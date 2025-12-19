package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

const (
	ContentTypeAppJson string = "application/json"
	ContentTypeAppForm string = "application/x-www-form-urlencoded"
)

func ResponseString(ctx *gin.Context, code int, message string) {
	ctx.String(code, message)
	ctx.Set("status_message", message)
}

func ResponseJSON(ctx *gin.Context, code int, obj any) {
	ctx.JSON(code, obj)
	o, _ := json.Marshal(obj)
	ctx.Set("status_message", string(o))
}
