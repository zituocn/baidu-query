package router

import (
	"git.zituo.net/zituocn/baidu-query/handler"
	"github.com/zituocn/gow"
)

func APIRouter(r *gow.Engine) {
	r.NoRoute(handler.Error404)
	v1 := r.Group("/v1")
	{
		v1.GET("/baidu", handler.BaiduQuery)
	}

}
