package handler

import (
	"fmt"
	"git.zituo.net/zituocn/baidu-query/util"
	"github.com/tidwall/gjson"
	"github.com/zituocn/gow"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	timeOut = 10
)

// Result 返回的结构体
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// BaiduQuery baidu 检测
func BaiduQuery(c *gow.Context) {
	s := c.GetString("url")
	if s == "" {
		c.JSON(&Result{
			Code: 1,
			Msg:  "请传入url参数",
		})
		return
	}

	flag, err := request(s)
	if err != nil {
		c.JSON(&Result{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(&Result{
		Code: 0,
		Msg:  "success",
		Data: struct {
			Url  string `json:"url"`
			Flag bool   `json:"flag"`
		}{
			Url:  s,
			Flag: flag,
		},
	})
}

func request(s string) (flag bool, err error) {
	//验证地址
	u, err := url.ParseRequestURI(s)
	if err != nil {
		err = fmt.Errorf("错误的url格式,如：https://www.baidu.com")
		return
	}
	uu := fmt.Sprintf("wd=%s&rn=1&tn=json&ie=utf-8&cl=3&f=9", u.String())

	//创建request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.baidu.com/s?%s", uu), nil)
	//指定useragent
	req.Header.Set("User-Agent", util.GetUserAgent())
	if err != nil {
		return
	}
	client := http.Client{
		Timeout: timeOut * time.Second,
	}

	//开始请求
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp == nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("错误的响应码 :%d", resp.StatusCode)
		return
	}

	//处理响应
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("读取响应出错 :%v", err)
		return
	}
	g := gjson.ParseBytes(b)
	result := g.Get("feed.entry.#.url")
	if len(result.Array()) > 0 {
		flag = true
	}
	return
}
