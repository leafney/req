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

// ~~get请求 设置请求参数 发现无法为请求添加上参数~~
// ~~通过比对 X2 和 X3的结果，后来发现 请求的req和设置参数的req是不同的req，所以才会赋不上值~~
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

// 测试jsonp
func TestX4(t *testing.T) {
	url := "https://p.3.cn/prices/mgets?callback=jQuery7955233&type=1&area=1_72_2819_0&pdtk=&pduid=163434863&pdpin=&pin=null&pdbp=0&skuIds=J_7437788%2CJ_32953185394%2CJ_31976609348&ext=11100000&source=item-pc"
	// url := "https://p.3.cn/prices/mgets?callback=jQuery7955233&type=1&area=1_72_2819_0&pdtk=&pduid=163434863&pdpin=&pin=null&pdbp=0&skuIds=J_74378&ext=11100000&source=item-pc"

	rs := XSets{
		"Debug": true,
	}

	s, _ := XReq(url, rs)
	t.Log(s.String())

	// 解析jsonp为对象
	// var p []jdp
	// e := s.ToJSONFromJSONP(&p)
	// if e != nil {
	// 	t.Error(e)
	// } else {
	// 	t.Log(p)
	// }

	// 解析jsonp为json字符串
	n, err := s.ToJSONStrFromJSONP()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(n)
	}

}

type jdp struct {
	Id string `id`
	P  string `json:"p"`
}

// 测试代理
func TestX5(t *testing.T) {

	header := Header{
		"Content-Type":    "application/json;charset=utf-8",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept":          "*/*",
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	}

	// param := Param{
	// 	"pdtk":     "",
	// 	"pduid":    "163434863",
	// 	"skuIds":   "J_3319362,J_3693867",
	// 	"callback": "jQuery199463",
	// 	"area":     "1_72_2819_0",
	// 	"type":     "1",
	// 	"source":   "item-pc",
	// }
	SetProxyUrl("27.203.219.181:8060")
	Debug = true

	// malformed HTTP response "<html>"
	//
	// panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	// 	panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x60 pc=0x12a75fc]
	//

	// https://p.3.cn/prices/mgets
	r, err := Get("https://www.cnblogs.com/", header)
	if err != nil {
		t.Error(err)
	}
	t.Log(r.String())

}
