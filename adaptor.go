package youtube_go

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type InnerTubeAdaptor struct {
	context ClientContext
	session *http.Client
}

func NewInnerTubeAdaptor(context ClientContext, session *http.Client) *InnerTubeAdaptor {
	if session == nil {
		session = &http.Client{}
	}
	return &InnerTubeAdaptor{
		context: context,
		session: session,
	}
}

func (ita *InnerTubeAdaptor) buildRequest(endpoint string, params map[string]string, body map[string]interface{}) (*http.Request, error) {
	body = Contextualise(ita.context, body)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.BaseURL+strings.ToLower(endpoint), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// 设置基于上下文的请求头
	req.Header = ita.context.Headers()

	// 设置查询参数
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (ita *InnerTubeAdaptor) request(endpoint string, params map[string]string, body map[string]interface{}) (*http.Response, error) {
	req, err := ita.buildRequest(endpoint, params, body)
	/*
		fmt.Println("Method: ", req.Method)
		fmt.Println("URL: ", req.URL)
		fmt.Println("Request Headers: ")
		for k, v := range req.Header {
			fmt.Println(fmt.Sprintf("%s: ", k), v)
		}
		//fmt.Println("Header: ", req.Header)
		//fmt.Println("Cookies: ", req.Cookies())
		fmt.Println("UserAgent: ", req.UserAgent())
		//fmt.Println("Form: ", req.Form)
		//fmt.Println("PostForm", req.PostForm)
		fmt.Println("Body: ", req.Body)
	*/
	if err != nil {
		return nil, err
	}
	return ita.session.Do(req)
}

func (ita *InnerTubeAdaptor) Dispatch(endpoint string, params map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ita.request(endpoint, params, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	//fmt.Println(""resp.Body)
	if contentType != "" && !isJSONContentType(contentType) {
		return nil, fmt.Errorf("expected JSON response, got %q", contentType)
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var reader io.Reader
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzr, err := gzip.NewReader(bytes.NewReader(bodyResp))
		if err != nil {
			return nil, err
		}
		defer gzr.Close()
		reader = gzr
	} else {
		reader = bytes.NewReader(bodyResp)
	}

	var responseData map[string]interface{}
	if err := json.NewDecoder(reader).Decode(&responseData); err != nil {
		return nil, err
	}

	visitorData, ok := responseData["responseContext"].(map[string]interface{})["visitorData"].(string)
	if ok {
		ita.context.XGoogVisitorId = visitorData
	}

	if errorData, ok := responseData["error"]; ok {
		return nil, fmt.Errorf("request error: %v", errorData)
	}

	return responseData, nil
}

func isJSONContentType(contentType string) bool {
	return contentType == "application/json" || contentType == "application/json; charset=utf-8" || contentType == "application/json; charset=UTF-8"
}
