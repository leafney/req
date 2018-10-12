package req

/*
req 扩展
*/

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// *************************

// 自定义设置参数
type XSets map[string]interface{}

// *************************

// 请求基类
func XRequest(url string, xSets map[string]interface{}, v ...interface{}) string {

	header := make(http.Header)

	// 默认配置

	// 设置默认Content-Type
	header.Set("Content-Type", "application/json;charset=utf-8")

	// 设置默认User-Agent
	header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

	// 设置默认超时时间 默认10s
	// req.SetTimeout(10 * time.Second)

	// 请求类型
	method_Type := "GET"
	// 数据类型   default jsonp
	data_Type := "default"

	// 设置自定义配置参数
	for key, value := range xSets {
		// fmt.Println("key:", key, "value:", value)

		ky := strings.ToLower(key)
		switch ky {
		case "content-type":
			header.Set("Content-Type", value.(string))
		case "user-agent":
			header.Set("User-Agent", value.(string))
		case "time-out":
			SetTimeout(time.Duration(value.(int)) * time.Second)
		case "proxy":
			SetProxyUrl("http://my.proxy.com:23456")
		case "referer":
			header.Set("Referer", value.(string))
		case "debug":
			Debug = value.(bool)
		case "data-type":
			// 请求类型
			data_Type = strings.ToLower(value.(string))
		case "method":
			// 请求类型
			method_Type = strings.ToUpper(value.(string))
		default:
			fmt.Sprintf("the %s-%s not default settings", key, value)
			// 其他的参数，默认加入请求头中
			header.Set(key, value.(string))
		}
	}

	var r *Resp
	var err error

	// 根据请求类型，设置请求参数或请求体，发起请求
	switch method_Type {
	case "GET":
		// 处理请求参数
		fmt.Println("发起get请求")
		// fmt.Println(v)
		r, err = Get(url, v...)
	case "POST":
		r, err = Post(url, v...)
	case "PUT":
		r, err = Put(url, v...)
	case "DELETE":
		r, err = Delete(url, v...)

	}

	if err != nil {
		fmt.Printf("[ERROR]request error: [%s] \n", err)
		// Todo 请求出错后该如何处理
	}

	fmt.Println("...*ResponseData*...")

	// 判断StatusCode
	resp := r.Response()
	fmt.Println(resp.StatusCode)

	// 返回结果--字符串
	resp_str := r.String()
	// 如果是 get 且 jsonp 则自动处理返回数据中的callback
	if method_Type == "GET" && data_Type == "jsonp" {
		pat := `^[^(]*?\((.*)\)[^)]*$` // 匹配 jQuery7955233(); 的jsonp格式
		reg := regexp.MustCompile(pat)
		resp := reg.FindStringSubmatch(resp_str)
		if len(resp) == 2 {
			resp_str = resp[1]
		}
	}

	fmt.Println(resp_str)
	return resp_str
}

// *************************
