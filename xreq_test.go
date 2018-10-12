package req

/*

go test -v -run TestA
*/

import (
	// "github.com/Leafney/xhttp"
	"net/http"
	"testing"
)

// 发起get请求
func TestX1(t *testing.T) {
	url := "http://httpbin.org/ip"
	rs := XSets{
		// "Content-Type": "application/javascript",
		"Debug":  true,
		"Method": "GET",
	}
	s, _ := XReq(url, rs)
	t.Log(s)
}

// get请求 设置请求参数 发现无法为请求添加上参数
// 通过比对 X2 和 X3的结果，后来发现 请求的req和设置参数的req是不同的req，所以才会赋不上值
// 经过修改，将参数v ...interface{} 在Get() 中传参时是 `v...` 而不是 `v` ，这样就可以了
func TestX2(t *testing.T) {

	// https://p.3.cn/prices/get?type=1&skuid=J_5338456
	url := "https://p.3.cn/prices/get"
	rs := XSets{
		"Content-Type": "application/javascript",
		"Debug":        true,
		"Method":       "GET",
	}
	pa := Param{
		"type":  "1",
		"skuid": "J_5338456",
	}

	str, _ := XReq(url, rs, pa)
	t.Log(str)
}

// 测试使用原生的 req 方法 的get请求绑定参数
func TestX3(t *testing.T) {
	header := make(http.Header)
	header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

	p := Param{
		"type":  "1",
		"skuid": "J_5338456",
	}
	Debug = true
	Get("https://p.3.cn/prices/get", header, p)

}

// 测试req的扩展
func TestX4(t *testing.T) {
	// https://www.baidu.com/s?wd=hello
	p := Param{
		"wd":    "hello",
		"skuid": "J_5338456",
	}
	Debug = true
	XReq("https://www.baidu.com/s", p)
}
