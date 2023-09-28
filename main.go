package main

import (
	"git.zituo.net/zituocn/baidu-query/router"
	"github.com/zituocn/gow"
)

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	router.APIRouter(r)
	r.Run()
}
