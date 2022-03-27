package req

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

const (
	ResponseSettledFlag = "ResponseSettledFlag"
	ResponseFlag        = "ResponseFlag"
	// ResponseCacheFlag 用于控制请求结果是否记入缓存，非200请求默认不记入缓存
	ResponseCacheFlag = "ResponseCacheFlag"
	ResponseErrorFlag = "ResponseErrorFlag"
)

// SetResponse 设置返回体
func SetResponse(c *gin.Context, status int, resp render.Render) {
	_, found := c.Get(ResponseSettledFlag)
	if !found {
		c.Set(ResponseFlag, resp)
		c.Set(ResponseSettledFlag, true)
	}
	c.Render(status, resp)
}

// GetResponse return response
func GetResponse(c *gin.Context) render.Render {
	if value, exists := c.Get(ResponseFlag); exists {
		return value.(render.Render)
	}
	return nil
}

func JSON(c *gin.Context, status int, data interface{}) {
	SetResponse(c, status, render.JSON{Data: data})
	c.Abort()
}

// JSONSuccess 成功输出，标准输出格式
func JSONSuccess(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, render.JSON{Data: &gin.H{
		"error": 0,
		"msg":   "success",
		"data":  data,
		"cache": c.GetBool(ResponseCacheFlag),
	}})
}

// JSONError 错误输出，标准输出格式
func JSONError(c *gin.Context, status int, errCode int, err error, data interface{}) {
	c.Set(ResponseCacheFlag, false)
	c.Set(ResponseErrorFlag, err)
	JSON(c, status, render.JSON{Data: &gin.H{
		"error": errCode,
		"msg":   err.Error(),
		"data":  data,
		"cache": false,
	}})
}
