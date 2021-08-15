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
	"time"
)

var httpClient1000 = createHTTPClient(1000)

func main()  {
	fmt.Println("------------START------------")

	for m := 0; m < 30; m ++ {
		for i := 1; i < 100; i ++ {
			go get()
			go get()
			go get()
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("------------END------------")

	time.Sleep(10 * time.Second)

	fmt.Println("------------Main END------------")
}

var apiUrl = "https://xiaoming.tradplusad.com/api/v1_2/adconf?v=1_2&sdkv=5.4.6&x=7DF4E3A96B8C3548604371A50F00EEAD&idfa=8818DA59-BBDE-4509-967D-1435D58E89D6&did=8818DA59-BBDE-4509-967D-1435D58E89D6&ddid=0CC00914-2C69-48F8-8C1B-BF847100E802&m=com.oversea.miniworld&iso=VN&z=Asia/Ho_Chi_Minh&l=vi-VN&os=2&appid=2F94F88F6D32AB441554C7E66EF00AC7&app_ver=0.53.11&app_instime=1626603968&device_idfa=0CC00914-2C69-48F8-8C1B-BF847100E802&device_idfv=EAD335F3-EB48-4590-8FB5-791CBDB6E226&device_osv=13.3.1&device_type=1&device_make=Apple&device_model=iPhone&device_contype=2"

var bytesData = ""




func get() int {

	response, body, err := HttpGetRequest(apiUrl, httpClient1000)
	fmt.Println("------------ERR START------------")
	fmt.Println(err)
	fmt.Println("------------ERR END------------")
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

	//if response.Header.Get("Content-Encoding") == "gzip" {
	//	body, err = gzip.NewReader(response.Body)
	//	if err != nil {
	//		return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", err))
	//	}
	//}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}

func HttpGetRequest(apiUrl string, httpClient *http.Client)  (*http.Response, []byte, error) {
	response, err := httpClient.Get(apiUrl)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, body, nil
}
