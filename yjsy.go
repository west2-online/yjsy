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

func (s *Student) GetWithFields(url string, kvs map[string]string) (*html.Node, error) {
	resp, err := s.NewRequest().SetHeader("Referer", constants.YjsyReferer).SetQueryParams(kvs).Get(url)
	// todo:目前我还不确定回话过期的结果是啥
	// 会话过期：会直接重定向，但我们禁用了重定向，所以会有error
	if err != nil {
		return nil, errno.CookieError
	}

	return htmlquery.Parse(bytes.NewReader(resp.Body()))
}
