package http

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
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

func Get(url string, headerOptions ...HeaderOption) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return responseHandle(resp, err)
}

func GetImage(url string, headerOptions ...HeaderOption) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetRespHeader(url string) (http.Header, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Header, nil
}

func GetRespCookies(url string) ([]*http.Cookie, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Cookies(), nil
}

func responseHandle(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	return b, nil
}

func transformString(t transform.Transformer, s string) (string, error) {
	r := transform.NewReader(strings.NewReader(s), t)
	b, err := io.ReadAll(r)
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
