package yjsy

import (
	"bytes"
	"crypto/tls"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"github.com/west2-online/yjsy/constants"
	"github.com/west2-online/yjsy/errno"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func NewStudent() *Student {
	// Disable HTTP/2.0
	// Disable Redirect
	client := resty.New().SetTransport(&http.Transport{
		TLSNextProto:    make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}).SetRedirectPolicy(resty.NoRedirectPolicy())

	return &Student{
		client: client,
	}
}
func (s *Student) WithLoginData(cookies []*http.Cookie) *Student {
	s.cookies = cookies
	s.client.SetCookies(cookies)
	return s
}

// WithUser 携带账号密码
func (s *Student) WithUser(id, password string) *Student {
	s.ID = id
	s.Password = password
	return s
}

func (s *Student) SetCookies(cookies []*http.Cookie) {
	s.cookies = cookies
	s.client.SetCookies(cookies)
}

func (s *Student) ClearLoginData() {
	s.cookies = []*http.Cookie{}
	s.client.Cookies = []*http.Cookie{}
}
func (s *Student) NewRequest() *resty.Request {
	return s.client.R()
}

func (s *Student) GetWithIdentifier(url string, queryParams map[string]string) (*html.Node, error) {
	request := s.NewRequest().SetHeader("Referer", constants.YJSYReferer)
	if queryParams != nil {
		for key, value := range queryParams {
			request = request.SetQueryParam(key, value)
		}
	}
	// 会话过期：会直接重定向，但我们禁用了重定向，所以会有error
	resp, err := request.Get(url)
	if err != nil {
		return nil, errno.CookieError
	}

	if strings.Contains(string(resp.Body()), "重新登录") {
		return nil, errno.CookieError
	}

	return htmlquery.Parse(bytes.NewReader(resp.Body()))
}

func (s *Student) PostWithIdentifier(url string, formData map[string]string) (*html.Node, error) {
	resp, err := s.NewRequest().SetHeader("Referer", constants.YJSYReferer).SetFormData(formData).Post(url)

	// 会话过期：会直接重定向，但我们禁用了重定向，所以会有error
	if err != nil {
		return nil, errno.CookieError.WithErr(err)
	}

	// id 或 cookie 缺失或者解析错误 TODO: 判断条件有点简陋
	if strings.Contains(string(resp.Body()), "处理URL失败") {
		return nil, errno.CookieError
	}

	return htmlquery.Parse(strings.NewReader(strings.TrimSpace(string(resp.Body()))))
}

func (s *Student) GetWithFields(url string, kvs map[string]string) (*html.Node, error) {
	resp, err := s.NewRequest().SetHeader("Referer", constants.YjsyCourseReferer).SetQueryParams(kvs).Get(url)
	// todo:目前我还不确定回话过期的结果是啥
	// 会话过期：会直接重定向，但我们禁用了重定向，所以会有error
	if err != nil {
		return nil, errno.CookieError
	}

	return htmlquery.Parse(bytes.NewReader(resp.Body()))
}
