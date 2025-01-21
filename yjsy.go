package yjsy

import (
	"crypto/tls"
	"net/http"

	"github.com/go-resty/resty/v2"
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
