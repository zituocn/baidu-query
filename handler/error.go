package handler

import "github.com/zituocn/gow"

// Error404 错误页面提示
func Error404(c *gow.Context) {
	c.ServerJSON(403, gow.H{
		"code": 1,
		"msg":  "HTTP 403 Forbidden",
	})
}
