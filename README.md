# baidu-query

baidu 收录检测

### 请使用新版本的baidu收录检测实现：

#### 项目地址：

```sh
https://github.com/zituocn/rich-api
```


### 安装

```sh
go get -u github.com/zituocn/baidu-query
```

或

```sh
git clone github.com/zituocn/baidu-query
go build -o baidu-query
./baidu-query
```



### API

*GET*

```
/v1/baidu?url=https://github.com/tidwall/gjson
```

*正确响应*

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "url": "https://github.com/tidwall/gjson",
    "flag": true
  }
}
```
其中

```json
flag 表示是否已收录
```

*错误的响应*

```json
{
    "code": 1,
    "msg": "请传入url参数",
    "data": null
}
```

或

```json
{
    "code": 1,
    "msg": "错误的url格式,如：https://www.baidu.com",
    "data": null
}
```


### 第三方库
* web框架 [github.com/zituocn/gow](github.com/zituocn/gow)
* json处理 [github.com/tidwall/gjson](github.com/tidwall/gjson)
