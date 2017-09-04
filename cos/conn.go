package cos

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Conn http 请求类
type Conn struct {
	c    *http.Client
	conf *Conf
}

// Do 所有请求的入口
func (conn *Conn) Do(ctx context.Context, method, bucket, object string, params map[string]interface{}, headers map[string]string, body io.Reader) (*http.Response, error) {
	queryStr := getQueryStr(params)
	url := conn.buildURL(bucket, object, queryStr)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	fmt.Println(url)
	conn.signHeader(req, params, headers)
	req.Header.Set("User-Agent", conn.conf.UA)
	req.Header.Set("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	setHeader(req, headers)

	res, err := conn.c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return checkHTTPErr(res)
	}

	return res, nil
}

func getQueryStr(params map[string]interface{}) string {
	if params == nil || len(params) == 0 {
		return ""
	}

	buf := new(bytes.Buffer)
	buf.WriteString("?")
	for k, v := range params {
		buf.WriteString(k)
		vs := interfaceToString(v)
		if vs == "" {
			buf.WriteString("&")
			continue
		}
		buf.WriteString("=")
		buf.WriteString(vs)
		buf.WriteString("&")
	}

	return strings.Trim(buf.String(), "&")
}

func (conn *Conn) buildURL(bucket, object, queryStr string) string {
	domain := fmt.Sprintf("%s-%s.cos.%s.%s", bucket, conn.conf.AppID, conn.conf.Region, conn.conf.Domain)
	url := fmt.Sprintf("http://%s/%s%s", domain, object, queryStr)

	return url
}

func setHeader(req *http.Request, headers map[string]string) {
	if headers == nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func checkHTTPErr(res *http.Response) (*http.Response, error) {
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return res, nil
	}

	err := HTTPError{}
	err.Code = res.StatusCode
	if res.StatusCode >= 300 && res.StatusCode < 400 {
		err.Message = "资源被重定向"
	}

	if res.StatusCode >= 400 && res.StatusCode < 500 {
		err.Message = "请求被拒绝"
	}

	if res.StatusCode >= 500 {
		err.Message = "cos服务器错误"
	}

	if res.ContentLength > 0 {
		resErr := &Error{}
		e := XMLDecode(res.Body, resErr)
		if e != nil {
			return nil, err
		}
		err.Message += resErr.Message
	}

	return res, err
}

// XMLDecode xml解析方法
func XMLDecode(r io.Reader, i interface{}) error {
	jd := xml.NewDecoder(r)
	err := jd.Decode(i)

	return err
}
