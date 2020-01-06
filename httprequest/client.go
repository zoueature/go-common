/*
 + ------------------------------------------------+
 | Author: Zoueature                               |
 + ------------------------------------------------+
 | Email: zoueature@gmail.com                      |
 + ------------------------------------------------+
 | Date: 2019/12/30                                |
 + ------------------------------------------------+
 | Time: 10:17                                     |
 + ------------------------------------------------+
 | Description:                                    |
 + ------------------------------------------------+
*/

package httprequest

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	*http.Client
}

type Response struct {
	Err error
	StatusCode int
	Body string
}

func NewClient() *Client {
	client := &http.Client{}
	localClient := &Client{client}
	return localClient
}

func (c *Client) Post(url string, params map[string]string, header http.Header) *Response {
	var body io.Reader
	if len(params) == 0 {
		body = nil
	} else {
		var paramStr string
		for key, value := range params {
			paramStr += key + "=" + value + "&"
		}
		paramStr = paramStr[:len(paramStr)-1]
		body = strings.NewReader(paramStr)
	}
	response := new(Response)
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		response.Err = err
		return response
	}
	if header != nil {
		request.Header = header
	}
	result, err := c.Client.Do(request)
	if err != nil {
		response.Err = err
		return response
	}
	response.StatusCode = result.StatusCode
	resultString, err := ioutil.ReadAll(result.Body)
	if err != nil {
		response.Err = err
		return response
	}
	response.Body = string(resultString)
	return response
}

func (c *Client) Get(requestUrl string, queries map[string]string, header http.Header) *Response {
	if queries != nil {
		p := url.Values{}
		for key, value := range queries {
			p.Add(key, value)
		}
		requestUrl += "?" + p.Encode()
	}
	response := new(Response)
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		response.Err = err
		return response
	}
	result, err := c.Client.Do(request)
	if err != nil {
		response.Err = err
		return response
	}
	response.StatusCode = result.StatusCode
	resultStr, err := ioutil.ReadAll(result.Body)
	if err != nil {
		response.Err = err
		return response
	}
	response.Body = string(resultStr)
	return response
}

func (c *Client) Request(method, requestUrl string, body io.Reader, header http.Header) (*http.Response, error) {
	newRequest, err := http.NewRequest(method, requestUrl, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		newRequest.Header = header
	}
	response, err := c.Do(newRequest)
	if err != nil {
		return nil, err
	}
	return response, err
}
