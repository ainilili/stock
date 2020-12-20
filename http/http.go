/**
2 * @Author: Nico
3 * @Date: 2020/12/20 20:27
4 */
package http

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var(
	httpClient = &http.Client{}
)

type HeaderOption struct {
	Name  string
	Value string
}

type QueryParameter struct {
	Key   string
	Value interface{}
}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}
}

func Get(url string, headerOptions ...HeaderOption) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func() {
		if resp != nil {
			if e := resp.Body.Close(); e != nil {
				fmt.Println(e)
			}
		}
	}()
	return responseHandle(resp, err)
}

func responseHandle(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respBody, _ := Decode(string(b), "gb18030")
	return respBody, nil
}

func transformString(t transform.Transformer, s string) (string, error) {
	r := transform.NewReader(strings.NewReader(s), t)
	b, err := ioutil.ReadAll(r)
	return string(b), err
}

func Decode(otherEncodeStr, encodeType string) (string, error) {
	e, _ := charset.Lookup(encodeType)
	if e == nil {
		return "", fmt.Errorf("%s: not found", encodeType)
	}
	decodeStr, err := transformString(e.NewDecoder(), otherEncodeStr)
	if err != nil {
		return "", err
	}
	return decodeStr, nil
}

func Encode(utf8EncodeStr, encodeType string) (string, error) {
	e, _ := charset.Lookup(encodeType)
	if e == nil {
		return "", fmt.Errorf("%s: not found", encodeType)
	}
	encodeStr, err := transformString(e.NewEncoder(), utf8EncodeStr)
	if err != nil {
		return "", err
	}
	return encodeStr, nil
}