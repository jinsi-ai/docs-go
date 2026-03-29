package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MsgResponse 消息响应
type MsgResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// Success 返回成功响应
func Success(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, MsgResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error 返回错误响应
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, MsgResponse{
		Success: false,
		Message: message,
	})
}

// BadRequest 返回400错误响应
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

// NotFound 返回404错误响应
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

// Unauthorized 返回401错误响应
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

// InternalServerError 返回500错误响应
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}
