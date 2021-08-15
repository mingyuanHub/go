package main

import (
	"compress/gzip"
	"io/ioutil"
	"net"
	"net/http"
	"bytes"
	"io"
	"errors"
	"fmt"
	"runtime"
	"time"
)

var httpClient1000 = createHTTPClient(1000)

func main()  {
	fmt.Println("------------START------------")

	for m := 0; m < 30; m ++ {
		for i := 1; i < 20; i ++ {
			go post()
			go post()
			go post()
		}
		fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}


	fmt.Println("------------END------------")

	time.Sleep(5 * time.Second)

	fmt.Println("------------Main END------------")
}

var apiUrl = "http://172.16.0.84/test/8.php"

var bytesData = ""




func post() int {

	var header = make(map[string]string)

	response, body, err := HttpPostRequest(apiUrl, []byte(bytesData), header, httpClient1000)
	//fmt.Println("------------ERR START------------")
	//fmt.Println(err)
	//fmt.Println("------------ERR END------------")
	return 0
	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Println(response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println(string(body))
		return 1
	}

	return 0
}


func createHTTPClient(requestTimeout int) *http.Client {
	transport := &http.Transport{
		MaxIdleConns: 1000,
		MaxIdleConnsPerHost: 1000,
		IdleConnTimeout:       300 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout: time.Duration(requestTimeout) * time.Millisecond,
	}
	return client
}

func HttpPostRequest(apiUrl string, bytesData []byte, headers map[string]string, httpClient *http.Client)  (*http.Response, []byte, error){
	var err error

	var isGzip = false
	if len(headers) > 0 {
		for key, item := range headers {
			if key == "Accept-Encoding" && item == "gzip" {
				isGzip = true
				break
			}
		}
	}

	var reader *bytes.Buffer

	if isGzip {
		var zBuf bytes.Buffer
		zw := gzip.NewWriter(&zBuf)
		if _, err = zw.Write(bytesData); err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("gzip error='%s'", err))
		}
		zw.Close()
		reader = &zBuf
	} else {
		reader = bytes.NewBuffer(bytesData)

	}

	request, err := http.NewRequest("POST", apiUrl, reader)

	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("newRequest error='%s'", err))
	}

	if len(headers) > 0 {
		for key, item := range headers {
			request.Header.Set(key, item)
		}
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}

	defer response.Body.Close()

	body := response.Body

	if response.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", err))
		}
	}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}
