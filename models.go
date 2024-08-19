package youtube_go

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Locale struct {
	hl string
	gl *string
}

func (l *Locale) AcceptLanguage() string {
	langs := []string{l.hl}
	if l.gl != nil {
		langs = append(langs, *l.gl)
	}
	return strings.Join(langs, ",")
}

type Error struct {
	code    int
	message string
	reason  string
}

func (e *Error) String() string {
	return fmt.Sprintf("%d %s: %s", e.code, http.StatusText(e.code), e.message)
}

func (e *Error) Code() int {
	return e.code
}

type ClientContext struct {
	ClientName     string
	ClientVersion  string
	ClientID       int
	APIKey         string
	UserAgent      string
	Referer        string
	Locale         *Locale
	XGoogVisitorId string
}

func (c *ClientContext) Params() map[string]string {
	params := map[string]string{
		"alt": "json",
	}
	if c.APIKey != "" {
		params["key"] = c.APIKey
	}
	return params
}

func (c *ClientContext) Context() map[string]string {
	return map[string]string{
		"clientName":    c.ClientName,
		"clientVersion": c.ClientVersion,
	}
}

func (c *ClientContext) Headers() http.Header {
	// 将 map[string]string 转换为 http.Header
	headers2 := http.Header{}
	headers2.Add("Host", config.Host)
	headers2.Add("Accept", "*/*")
	headers2.Add("Accept-Encoding", "gzip, deflate")
	headers2.Add("Connection", "keep-alive")
	headers2.Add("Content-Type", "application/json")
	headers2.Add("X-Goog-Api-Format-Version", "1")
	headers2.Add("X-YouTube-Client-Name", fmt.Sprintf("%s", strconv.Itoa(c.ClientID)))
	headers2.Add("X-YouTube-Client-Version", c.ClientVersion)

	headers := map[string]string{
		"X-Goog-Api-Format-Version": "1",
		"X-YouTube-Client-Name":     fmt.Sprintf("%s", strconv.Itoa(c.ClientID)),
		"X-YouTube-Client-Version":  c.ClientVersion,
	}
	if c.UserAgent != "" {
		headers["User-Agent"] = c.UserAgent
		headers2.Add("User-Agent", c.UserAgent)
	}
	if c.Referer != "" {
		headers["Referer"] = c.Referer
		headers2.Add("Referer", c.Referer)
	}
	if c.Locale != nil {
		headers["Accept-Language"] = c.Locale.AcceptLanguage()
		headers2.Add("Accept-Language", c.Locale.AcceptLanguage())
	}

	if c.XGoogVisitorId != "" {
		headers2.Add("X-Goog-Visitor-Id", c.XGoogVisitorId)
	}
	return headers2
}

type Config struct {
	Host    string
	BaseURL string
	Clients []ClientContext
}

type ResponseContext struct {
	Function    *string
	BrowseID    *string
	Context     *string
	VisitorData *string
	Client      Client
	Request     Request
	Flags       Flags
}

type Request struct {
	Type *string
	ID   *string
}

type Client struct {
	Name    *string
	Version *string
}

type Flags struct {
	LoggedIn *bool
}

type ResponseFingerprint struct {
	Request  *string
	Function *string
	BrowseID *string
	Context  *string
	Client   *string
}
